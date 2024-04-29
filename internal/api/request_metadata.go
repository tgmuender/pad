package api

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
	"strings"
	storage "xgmdr.com/pad/internal/storage/model"
)

type RequestMetadata struct {
	headerValues []string
}

func extractAuthentication(context context.Context) *AuthenticatedIdentity {
	md, _ := metadata.FromIncomingContext(context)

	authorization := md.Get("authorization")
	fmt.Printf("authorization: %s", authorization)

	requestMetadata := RequestMetadata{headerValues: md.Get("authorization")}
	fmt.Println(requestMetadata)

	return requestMetadata.extractAuthentication()
}

func (m *RequestMetadata) extractAuthentication() *AuthenticatedIdentity {
	var idToken string
	var prefix = "Bearer "
	for _, headerValue := range m.headerValues {
		fmt.Println(headerValue)
		if strings.HasPrefix(headerValue, prefix) {
			fmt.Println(headerValue)
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

type AuthenticatedIdentity struct {
	Issuer  string
	Email   string
	Subject string
}

func (identity *AuthenticatedIdentity) toOwner() *storage.Owner {
	return &storage.Owner{
		Issuer:  identity.Issuer,
		OwnerId: identity.Subject,
		Email:   identity.Email,
	}
}
