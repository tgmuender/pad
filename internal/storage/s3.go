package storage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	url "net/url"
	"strings"
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

func (s *S3StorageService) GetPreSignedUrl(objectName, method string) (string, error) {
	ctx := context.Background()
	duration, _ := time.ParseDuration("1h") // Example duration, adjust as needed

	var presignedURL *url.URL
	var err error

	switch strings.ToUpper(method) {
	case "GET":
		presignedURL, err = s.client.PresignedGetObject(ctx, s.bucketName, objectName, duration, nil)
	case "PUT":
		presignedURL, err = s.client.PresignedPutObject(ctx, s.bucketName, objectName, duration)
	default:
		return "", fmt.Errorf("unsupported method: %s", method)
	}

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
