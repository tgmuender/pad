package api

import (
	pb "xgmdr.com/pad/proto"
)

type PetApi struct {
	pb.UnimplementedPetServiceServer
}
