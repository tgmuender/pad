package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	storage "xgmdr.com/pad/internal/storage/model"
	pb "xgmdr.com/pad/proto"
)

func (m *PetApi) NewPet(grpcContext context.Context, request *pb.NewPetRequest) (*pb.NewPetResponse, error) {
	fmt.Println("New Pet request received: ", request.String())

	authentication := extractAuthentication(grpcContext)

	fmt.Println("Subject: ", authentication.Subject)
	fmt.Println("email: ", authentication.Email)

	petEntity, _ := toPetEntity(authentication, request)

	if err := crdbgorm.ExecuteTx(context.Background(), storage.Db, nil,
		func(tx *gorm.DB) error {
			return storage.InsertPet(petEntity)
		},
	); err != nil {
		// For information and reference documentation, see:
		//   https://www.cockroachlabs.com/docs/stable/error-handling-and-troubleshooting.html
		fmt.Println(err)
	}

	return &pb.NewPetResponse{
		Id:   petEntity.ID.String(),
		Name: request.GetName(),
	}, nil
}

func toPetEntity(identity *AuthenticatedIdentity, request *pb.NewPetRequest) (*storage.PetEntity, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("Unable to generate id for new pet request: %v", request.Name)
		return nil, err
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return &storage.PetEntity{
		ID:    id,
		Owner: *identity.toOwner(),
		Data:  string(requestData),
	}, nil
}
