package api

import (
	"context"
	"encoding/json"
	"fmt"
	storage "xgmdr.com/pad/internal/storage/model"
	pb "xgmdr.com/pad/proto"
)

func (m *PetApi) ListPets(grpcContext context.Context, request *pb.ListPetsRequest) (*pb.ListPetsResponse, error) {
	fmt.Println("list pet request received: ", request.String())

	authentication := extractAuthentication(grpcContext)
	owner := authentication.toOwner()

	var result []storage.PetEntity
	storage.Db.Where(
		&storage.PetEntity{
			Owner: storage.Owner{
				Issuer:  owner.Issuer,
				OwnerId: owner.OwnerId,
			},
		},
	).Find(&result)

	fmt.Printf("Found %d entities", len(result))

	var pets []*pb.Pet
	for idx, pe := range result {
		fmt.Printf("Processing item %d", idx)
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
