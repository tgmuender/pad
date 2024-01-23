package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"xgmdr.com/pad/internal/api"
	pb "xgmdr.com/pad/proto"
)

func newServerCmd() *cobra.Command {
	command := cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetInt("port")
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			s := grpc.NewServer()

			pb.RegisterPetServiceServer(s, &api.PetApi{})
			log.Printf("server listening at %v", lis.Addr())
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		},
	}

	command.PersistentFlags().Int("port", 8000, "Port to listen for incoming requests.")

	return &command
}
