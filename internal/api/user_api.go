package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "xgmdr.com/pad/proto"
)

type UserAPi struct {
	pb.UnimplementedUserServiceServer
}

func (m *UserAPi) WhoAmI(grpcContext context.Context, request *emptypb.Empty) (*pb.UserResponse, error) {
	return &pb.UserResponse{Email: "asdf@xgmdr.com"}, nil
}
