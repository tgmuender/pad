package api

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
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
}

type RequestMetadata struct {
	headerValues []string
}

func ExtractAuthentication(context context.Context) *AuthenticatedIdentity {
	md, _ := metadata.FromIncomingContext(context)

	requestMetadata := RequestMetadata{headerValues: md.Get("authorization")}

	return requestMetadata.extractAuthentication()
}

func (m *RequestMetadata) extractAuthentication() *AuthenticatedIdentity {
	var idToken string
	var prefix = "Bearer "
	for _, headerValue := range m.headerValues {
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

	introspectToken(token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		var subject, _ = claims.GetSubject()
		var issuer, _ = claims.GetIssuer()

		return &AuthenticatedIdentity{
			Email:   fmt.Sprintf("%s", claims["email"]),
			Subject: subject,
			Issuer:  issuer,
		}
	} else {
		fmt.Println(err)
	}

	return nil
}

func (identity *AuthenticatedIdentity) ToOwner() *storage.Owner {
	return &storage.Owner{
		Issuer:  identity.Issuer,
		OwnerId: identity.Subject,
		Email:   identity.Email,
	}
}
