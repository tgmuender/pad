package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
	"xgmdr.com/pad/proto"
)

// Global variable with type: "pointer to cobra command
var newPetCommand *cobra.Command = &cobra.Command{
	Use:  "new",
	Long: `Creates a new pet`,
	Run: func(cmd *cobra.Command, args []string) {
		var defaultMsg = "TODO: create new pet"
		fmt.Fprintln(os.Stdout, defaultMsg)
		// Do Stuff Here

		var serverAddress = "localhost:8000"

		dial, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Fprint(os.Stderr, "error while connecting to server", err)
			os.Exit(1)
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		client := proto.NewPetServiceClient(dial)

		name := "Mangoloni"

		client.NewPet(ctx, &proto.NewPetRequest{
			Name: &name,
		})

	},
}

func main() {
	if err := newPetCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
