package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
	"xgmdr.com/pad/internal/logger"
	"xgmdr.com/pad/internal/storage"
)

// AuthenticatedIdentity Represents the authentication of the request. Should provide enough information to uniquely extract the required
// user information. Although (issuer, subject) is technically enough to uniquely identify the user, the email  address
// is added just to make identification easier when the data is screened by a human.
type AuthenticatedIdentity struct {
	// The OIDC issuer which issued the token.
	Issuer string

	// The subject of the token.
	Subject string

	// The email address from the token.
	Email string

	// The name of the user.
	Name string
}

func (identity *AuthenticatedIdentity) ToUser() *storage.User {
	return &storage.User{
		Issuer:  identity.Issuer,
		Subject: identity.Subject,
		Email:   identity.Email,
		Name:    identity.Name,
	}
}

// AuthenticatingInterceptor is a gRPC interceptor that authenticates the incoming request.
func AuthenticatingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Extract the authentication information from the request context.
		requestMetadata, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "no metadata found in request")
		}

		authHeader, ok := requestMetadata["authorization"]
		if !ok || len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authorization token not found in request metadata")
		}

		identity := extractIdentity(authHeader)

		user, err := GetOrCreateUser(identity)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unable to authenticate user")
		}

		ctx = context.WithValue(ctx, "user", user)

		return handler(ctx, req)
	}
}

// GetOrCreateUser queries the database to find a user with the given email address. If the user does not exist, it creates
// a new user.
func GetOrCreateUser(authenticated *AuthenticatedIdentity) (*storage.User, error) {
	if user, _ := GetUserByEmail(authenticated); user != nil {
		return user, nil
	}

	newUser := authenticated.ToUser()
	if err := storage.InsertUser(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

// GetUserByEmail queries the database to find a user with the given email address.
func GetUserByEmail(authenticated *AuthenticatedIdentity) (*storage.User, error) {
	if authenticated == nil {
		return nil, errors.New("authenticated identity must not be nil")
	}

	user, err := storage.FindUserByEmail(authenticated.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// extractIdentity extracts the identity from the authorization request headers.
func extractIdentity(headers []string) *AuthenticatedIdentity {
	var idToken string
	var prefix = "Bearer "
	for _, headerValue := range headers {
		logger.Get().Debug(
			"Found 'Authorization' header",
			zap.String("header", headerValue),
		)

		if strings.HasPrefix(headerValue, prefix) {
			logger.Get().Debug(
				"Found bearer token",
				zap.String("token", headerValue),
			)
			idToken = headerValue
		}
	}

	if idToken == "" {
		// no id token found in cookies
		return nil
	}

	idToken = idToken[len(prefix):len(idToken)]

	token, err := jwt.Parse(
		idToken,
		func(token *jwt.Token) (interface{}, error) {
			// Already verified at relying party
			return "ok", nil
		},
	)

	printToken(token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		var subject, _ = claims.GetSubject()
		var issuer, _ = claims.GetIssuer()

		return &AuthenticatedIdentity{
			Email:   fmt.Sprintf("%s", claims["email"]),
			Subject: subject,
			Issuer:  issuer,
			Name:    fmt.Sprintf("%s", claims["name"]),
		}
	} else {
		fmt.Println(err)
	}

	return nil
}
