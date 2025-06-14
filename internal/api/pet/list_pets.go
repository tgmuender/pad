package pet

import (
	"context"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"xgmdr.com/pad/internal/logger"
	"xgmdr.com/pad/internal/storage"
	pb "xgmdr.com/pad/proto"
)

// ListPets finds all pets having their owner set to the authenticated party in the 'grpcContext'.
func (m *Api) ListPets(grpcContext context.Context, request *pb.ListPetsRequest) (*pb.ListPetsResponse, error) {
	logger.Get().Debug("List pets request received")

	user, ok := grpcContext.Value("user").(*storage.User)
	if !ok {
		logger.Get().Debug("User not found")
		return nil, errors.New("user not found in context")
	}

	var result = storage.FindByOwner(user)

	logger.Get().Debug("Pet count matching user ", zap.Int("count", len(result)))

	var pets []*pb.Pet
	for idx, pe := range result {
		logger.Get().Debug("Processing item", zap.Int("idx", idx))

		var npr pb.NewPetRequest
		_ = json.Unmarshal([]byte(pe.Data), &npr)

		// Check if the pet has a profile picture
		var url string
		if pe.ProfilePicture != nil {
			url, _ = m.S3.GetPreSignedUrl(pe.ProfilePicture.ObjectKey, "get")
		}

		pets = append(pets, &pb.Pet{
			Id:                pe.ID.String(),
			Name:              npr.Name,
			Type:              npr.Type,
			ProfilePictureUrl: url,
		})
	}

	return &pb.ListPetsResponse{Pets: pets}, nil
}
