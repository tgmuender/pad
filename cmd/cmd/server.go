package cmd

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	credentials "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"strings"
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

			minioClient, err := initMinioClient(cmd.Flags())
			if err != nil {
				log.Fatalf("failed to initialize MinIO client: %v", err)
			}

			service := storage.NewS3StorageService(minioClient, "petadvisor")
			if err := service.PrepareBucket(); err != nil {
				log.Fatalf("failed to prepare bucket: %v", err)
			}
			url, err := service.GetPreSignedUrl("/pets/12345/image.jpg")

			fmt.Println(url)

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

// initMinioClient initializes the MinIO client with the provided flags.
func initMinioClient(flags *pflag.FlagSet) (*minio.Client, error) {
	s3endpoint, _ := flags.GetString("s3-endpoint")
	accessKey, _ := flags.GetString("s3-accessKeyId")
	secretKey, _ := flags.GetString("s3-secretAccessKey")

	return minio.New(s3endpoint, &minio.Options{
		Secure: strings.HasPrefix(s3endpoint, "https://"),
		Creds: credentials.NewStaticV4(
			accessKey,
			secretKey,
			"",
		),
	},
	)
}
