package pet

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"xgmdr.com/pad/internal/logger"
	"xgmdr.com/pad/internal/storage"
	pb "xgmdr.com/pad/proto"
)

// SetProfilePicture creates a pre-signed URL for uploading a pet's profile picture.
func (m *Api) SetProfilePicture(grpcContext context.Context, request *pb.SetProfilePictureRequest) (*pb.SetProfilePictureResponse, error) {
	logger.Get().Debug("Starting set profile picture request")

	user, ok := grpcContext.Value("user").(*storage.User)
	if !ok {
		logger.Get().Debug("User not found", zap.String("email", user.Email))
		return nil, errors.New("user not found in context")
	}

	//TODO: check access rights

	objectName := "pets/" + request.PetId + "/" + request.Filename // Example object name, adjust as needed
	if url, err := m.S3.GetPreSignedUrl(objectName, "put"); err != nil {
		logger.Get().Error("Failed to get pre-signed URL", zap.Error(err))
		return nil, err
	} else {
		logger.Get().Debug("Pre-signed URL generated", zap.String("url", url))

		fm := &storage.FileMetadata{
			ObjectKey:  objectName,
			PetID:      uuid.MustParse(request.PetId),
			Name:       request.Filename,
			UploaderID: user.Id,
			Type:       storage.ProfilePicture,
		}

		if err := storage.InsertFileMetadata(grpcContext, fm); err != nil {
			logger.Get().Error("Failed to insert file metadata", zap.Error(err))
			return nil, err
		}

		return &pb.SetProfilePictureResponse{
			PetId:     "Picture set successfully",
			UploadUrl: url,
		}, nil
	}
}
