package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"xgmdr.com/pad/internal/client"
	"xgmdr.com/pad/proto"
)

// deletePetCommand creates a new command to delete a pet.
func newDeletePetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a pet",
		Run: func(cmd *cobra.Command, args []string) {
			endpoint, _ := cmd.Flags().GetString("endpoint")
			id, _ := cmd.Flags().GetString("id")

			apiClient, ctx, cancel := client.PetServiceGrpcClient(endpoint)
			defer cancel()

			ctx, err := withAccessToken(ctx)
			if err != nil {
				fmt.Println("Error reading token:", err)
				os.Exit(1)
			}

			_, err = apiClient.DeletePet(ctx, &proto.DeletePetRequest{
				Id: id,
			})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		},
	}

	cmd.Flags().String("id", "", "ID of the pet to delete")
	cmd.MarkFlagRequired("id")

	return cmd
}
