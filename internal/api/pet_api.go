package api

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	pb "xgmdr.com/pad/proto"
)

type PetApi struct {
	pb.UnimplementedPetServiceServer
}

func (m PetApi) NewPet(context context.Context, request *pb.NewPetRequest) (*pb.NewPetResponse, error) {
	fmt.Print("New Pet request received: ", request.String())

	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("Unable to generate id for new pet request: %v", request.Name)
		return nil, err
	}
	return &pb.NewPetResponse{Id: id.String(), Name: request.GetName()}, nil
}
