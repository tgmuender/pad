package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
	"xgmdr.com/pad/proto"
)

func PetServiceGrpcClient(endpoint string) (proto.PetServiceClient, context.Context, context.CancelFunc) {
	dial, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Fprint(os.Stderr, "error while connecting to server", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return proto.NewPetServiceClient(dial), ctx, cancel
}
