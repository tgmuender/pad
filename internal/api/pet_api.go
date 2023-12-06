package api

import (
	"context"
	pb "xgmdr.com/pad/proto"
)

type PetApi struct {
	pb.UnimplementedPetServiceServer
}

func (m PetApi) NewPet(context.Context, *pb.NewPetRequest) (*pb.NewPetResponse, error) {
	return &pb.NewPetResponse{Id: "1", Name: "Mango"}, nil
}
