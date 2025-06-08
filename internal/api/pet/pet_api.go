package pet

import (
	"xgmdr.com/pad/internal/storage"
	pb "xgmdr.com/pad/proto"
)

type Api struct {
	pb.UnimplementedPetServiceServer
	S3 *storage.S3StorageService
}
