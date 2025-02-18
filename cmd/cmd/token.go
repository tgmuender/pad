package cmd

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
)

func withAccessToken(ctx context.Context) (context.Context, error) {
	token, err := readToken()
	if err != nil {
		return nil, err
	}

	md := metadata.New(map[string]string{"authorization": "Bearer " + token.AccessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx, nil
}

func readToken() (*oauth2.Token, error) {
	tokenData, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return nil, err
	}

	var token oauth2.Token
	if err := json.Unmarshal(tokenData, &token); err != nil {
		return nil, err
	}

	return &token, nil
}
