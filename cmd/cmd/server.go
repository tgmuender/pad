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
	"xgmdr.com/pad/internal/api/pet"
	"xgmdr.com/pad/internal/api/user"
	"xgmdr.com/pad/internal/storage"
	pb "xgmdr.com/pad/proto"
)

func newServerCmd() *cobra.Command {
	command := cobra.Command{
		Use:   "server",
		Short: "Starts the api server.",
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetInt("port")
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			s := grpc.NewServer(
				grpc.UnaryInterceptor(api.AuthenticatingInterceptor()),
			)

			pb.RegisterPetServiceServer(s, &pet.Api{})
			pb.RegisterUserServiceServer(s, &user.Api{})
			log.Printf("server listening at %v", lis.Addr())

			db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
			if err != nil {
				log.Fatal(err)
			}

			storage.Db = db

			db.AutoMigrate(
				&storage.User{},
				&storage.Pet{},
				&storage.MealEntity{},
			)

			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}

		},
	}

	command.PersistentFlags().Int("port", 8000, "Port to listen for incoming requests.")

	return &command
}
