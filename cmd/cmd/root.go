package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "padctl",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.AddCommand(newServerCmd())
	cmd.AddCommand(newPetCommand())
	cmd.AddCommand(newListPetsCommand())

	return &cmd
}
