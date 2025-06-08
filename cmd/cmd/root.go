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
	cmd.PersistentFlags().String("s3-endpoint", "localhost:9000", "S3 endpoint for storage communication.")
	cmd.PersistentFlags().String("s3-accessKeyId", "", "Access key id for S3.")
	cmd.PersistentFlags().String("s3-secretAccessKey", "", "Secret access key for S3.")
	cmd.PersistentFlags().String("s3-bucket-name", "petadvisor", "Bucket name for S3 storage.")

	return &cmd
}
