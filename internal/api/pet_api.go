package api

import (
	"context"
	"fmt"
	pb "xgmdr.com/pad/proto"
)

type PetApi struct {
	pb.UnimplementedPetServiceServer
}

func (m PetApi) NewPet(context context.Context, request *pb.NewPetRequest) (*pb.NewPetResponse, error) {
	fmt.Print("New Pet request received: ", request.String())
	return &pb.NewPetResponse{Id: "1", Name: "Mango"}, nil
}
