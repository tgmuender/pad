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
	pets map[string]string
}

func (m *PetApi) NewPet(context context.Context, request *pb.NewPetRequest) (*pb.NewPetResponse, error) {
	fmt.Println("New Pet request received: ", request.String())

	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("Unable to generate id for new pet request: %v", request.Name)
		return nil, err
	}

	if m.pets == nil {
		fmt.Println("Initializing pet map")
		m.pets = make(map[string]string)
	}
	m.pets[id.String()] = request.GetName()

	return &pb.NewPetResponse{Id: id.String(), Name: request.GetName()}, nil
}

func (m *PetApi) ListPets(context context.Context, request *pb.ListPetsRequest) (*pb.ListPetsResponse, error) {
	var pets []*pb.Pet
	for key, value := range m.pets {
		pets = append(pets, &pb.Pet{Id: key, Name: value})
	}
	return &pb.ListPetsResponse{Pets: pets}, nil
}
