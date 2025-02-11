package pet

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"xgmdr.com/pad/internal/logger"
	"xgmdr.com/pad/internal/storage"
	pb "xgmdr.com/pad/proto"
)

// NewPet creates and persists a new pet entity in the database.
// Returns a response containing the unique id of the created pet.
func (m *Api) NewPet(grpcContext context.Context, request *pb.NewPetRequest) (*pb.NewPetResponse, error) {
	logger.Get().Debug("New pet request received")

	user, ok := grpcContext.Value("user").(*storage.User)
	if !ok {
		logger.Get().Debug("User not found", zap.String("email", user.Email))
		return nil, errors.New("user not found in context")
	}

	petEntity, _ := toPetEntity(user, request)

	if err := storage.InsertPet(petEntity); err != nil {
		return nil, err
	}

	return &pb.NewPetResponse{
		Id:   petEntity.ID.String(),
		Name: request.GetName(),
	}, nil
}

// toPetEntity converts a NewPetRequest into a PetEntity.
func toPetEntity(user *storage.User, request *pb.NewPetRequest) (*storage.PetEntity, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("Unable to generate id for new pet request: %v", request.Name)
		return nil, err
	}

	requestData, err := protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
	}.Marshal(request)
	if err != nil {
		return nil, err
	}

	return &storage.PetEntity{
		ID:      id,
		OwnerID: user.Id,
		Data:    string(requestData),
	}, nil
}
