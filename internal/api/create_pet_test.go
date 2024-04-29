package api

import (
	"encoding/json"
	"fmt"
	"testing"
	storage "xgmdr.com/pad/internal/storage/model"
	pb "xgmdr.com/pad/proto"
)

func TestSerialization(t *testing.T) {
	petEntity := new(storage.PetEntity)

	p := new(pb.NewPetRequest)

	petEntity.Data = *p

	marshal, err := json.Marshal(petEntity)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
