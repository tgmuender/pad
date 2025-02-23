package pet

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"xgmdr.com/pad/internal/logger"
	"xgmdr.com/pad/internal/storage"
	"xgmdr.com/pad/proto"
)

// DeletePet deletes a pet.
func (m *Api) DeletePet(ctx context.Context, req *proto.DeletePetRequest) (*proto.DeletePetResponse, error) {
	logger.Get().Debug("Delete pet request received")

	user, ok := ctx.Value("user").(*storage.User)
	if !ok {
		logger.Get().Debug("User not found", zap.String("email", user.Email))
		return nil, errors.New("user not found in context")
	}

	petId, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	if !storage.ExistsForOwner(petId, user) {
		return nil, errors.New("pet not found")
	}

	if err := storage.DeletePet(petId); err != nil {
		return nil, err
	}

	return &proto.DeletePetResponse{}, nil
}
