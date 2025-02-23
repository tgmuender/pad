package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"xgmdr.com/pad/internal/client"
	"xgmdr.com/pad/proto"
)

func newPetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new pet",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("please provide a name")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			endpoint, _ := cmd.Flags().GetString("endpoint")
			petType, _ := cmd.Flags().GetString("type")
			sex, _ := cmd.Flags().GetString("sex")

			apiClient, ctx, cancel := client.PetServiceGrpcClient(endpoint)
			defer cancel()

			ctx, err := withAccessToken(ctx)
			if err != nil {
				fmt.Println("Error reading token:", err)
				os.Exit(1)
			}

			name := args[0]

			pet, err := apiClient.NewPet(ctx, &proto.NewPetRequest{
				Name: name,
				Type: petType,
				Sex:  mapSex(sex),
			})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(pet)
		},
	}

	cmd.Flags().String("type", "", "Type of pet (dog or cat)")
	cmd.Flags().String("sex", "", "male/female")

	return cmd
}

func mapSex(s string) proto.Sex {
	switch strings.ToLower(s) {
	case "male":
		return proto.Sex_MALE
	case "female":
		return proto.Sex_FEMALE
	default:
		return proto.Sex_UNKNOWN
	}
}
