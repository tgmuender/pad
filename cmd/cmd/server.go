package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"xgmdr.com/pad/internal/api"
	storage "xgmdr.com/pad/internal/storage/model"
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
			pb.RegisterUserServiceServer(s, &api.UserAPi{})
			log.Printf("server listening at %v", lis.Addr())

			db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
			if err != nil {
				log.Fatal(err)
			}

			storage.Db = db

			db.AutoMigrate(&storage.PetEntity{})

			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}

		},
	}

	command.PersistentFlags().Int("port", 8000, "Port to listen for incoming requests.")

	return &command
}
