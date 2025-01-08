package pet

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"xgmdr.com/pad/internal/api"
	"xgmdr.com/pad/internal/logger"
	"xgmdr.com/pad/internal/storage"
	pb "xgmdr.com/pad/proto"
)

// ListPets finds all pets having their owner set to the authenticated party in the 'grpcContext'.
func (m *Api) ListPets(grpcContext context.Context, request *pb.ListPetsRequest) (*pb.ListPetsResponse, error) {
	logger.Get().Debug("List pets request received")

	authentication := api.ExtractAuthentication(grpcContext)

	owner := authentication.ToOwner()

	var result = storage.FindByOwner(owner)

	logger.Get().Debug("Pet count matching owner ", zap.Int("count", len(result)))

	var pets []*pb.Pet
	for idx, pe := range result {
		logger.Get().Debug("Processing item", zap.Int("idx", idx))

		var npr pb.NewPetRequest
		_ = json.Unmarshal([]byte(pe.Data), &npr)
		pets = append(pets, &pb.Pet{
			Id:   pe.ID.String(),
			Name: npr.Name,
			Type: npr.Type,
		})
	}

	return &pb.ListPetsResponse{Pets: pets}, nil
}
