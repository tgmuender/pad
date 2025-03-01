package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	client "xgmdr.com/pad/internal/client"
	"xgmdr.com/pad/proto"
)

// newListPetsCommand creates a new command listing all pets to which the currently authenticated user has access to.
func newListPetsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all pets",
		Run: func(cmd *cobra.Command, args []string) {
			endpoint, _ := cmd.Flags().GetString("endpoint")

			apiClient, ctx, cancel := client.PetServiceGrpcClient(endpoint)
			defer cancel()

			ctx, err := withAccessToken(ctx)
			if err != nil {
				fmt.Println("Error reading token:", err)
				os.Exit(1)
			}

			res, err := apiClient.ListPets(ctx, &proto.ListPetsRequest{})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Printf("Found %v pets: ", len(res.Pets))
			fmt.Println(res.Pets)
		},
	}
	
	return cmd
}
