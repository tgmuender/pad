package cmd

import "github.com/spf13/cobra"

// NewRootCmd returns the root command for the application
func NewRootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "petctl",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(newServerCmd())
	cmd.AddCommand(newPetCommand())
	cmd.AddCommand(newListPetsCommand())
	cmd.AddCommand(newDeletePetCommand())
	cmd.AddCommand(loginCommand())

	cmd.PersistentFlags().String("endpoint", "localhost:8000", "Server endpoint for api communication.")

	return &cmd
}
