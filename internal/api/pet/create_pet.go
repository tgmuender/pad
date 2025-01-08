package pet

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"xgmdr.com/pad/internal/api"
	"xgmdr.com/pad/internal/logger"
	"xgmdr.com/pad/internal/storage"
	pb "xgmdr.com/pad/proto"
)

// NewPet creates and persists a new pet entity in the database.
// Returns a response containing the unique id of the created pet.
func (m *Api) NewPet(grpcContext context.Context, request *pb.NewPetRequest) (*pb.NewPetResponse, error) {
	logger.Get().Debug("New pet request received")

	authentication := api.ExtractAuthentication(grpcContext)

	petEntity, _ := toPetEntity(authentication, request)

	err := storage.InsertPet(petEntity)
	if err != nil {
		return nil, err
	}

	return &pb.NewPetResponse{
		Id:   petEntity.ID.String(),
		Name: request.GetName(),
	}, nil
}

func toPetEntity(identity *api.AuthenticatedIdentity, request *pb.NewPetRequest) (*storage.PetEntity, error) {
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
		ID:    id,
		Owner: *identity.ToOwner(),
		Data:  string(requestData),
	}, nil
}
