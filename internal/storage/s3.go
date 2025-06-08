package storage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"time"
)

// S3StorageService provides methods to interact with S3 storage.
type S3StorageService struct {
	client     *minio.Client
	bucketName string
}

// NewS3StorageService creates a new S3StorageService instance.
func NewS3StorageService(client *minio.Client, bucketName string) *S3StorageService {
	return &S3StorageService{
		client:     client,
		bucketName: bucketName,
	}
}

func (s *S3StorageService) GetPreSignedUrl(objectName string) (string, error) {
	ctx := context.Background()

	duration, _ := time.ParseDuration("1h") // Example duration, adjust as needed

	// Generate a pre-signed URL for the object
	presignedURL, err := s.client.PresignedPutObject(ctx, s.bucketName, objectName, duration)
	if err != nil {
		fmt.Println("Error generating pre-signed URL:", err)
		return "", err
	}

	return presignedURL.String(), nil
}

// PrepareBucket ensures that the bucket exists and is ready for use.
// It creates the bucket if it does not exist.
func (s *S3StorageService) PrepareBucket() error {
	ctx := context.Background()

	exists, err := s.client.BucketExists(ctx, s.bucketName)
	if err != nil {
		fmt.Println("Error checking if bucket exists:", err)
		return err
	}

	if !exists {
		err = s.client.MakeBucket(ctx, s.bucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Println("Error creating bucket:", err)
			return err
		}
	}

	fmt.Println("Bucket is ready:", s.bucketName)

	return nil
}
