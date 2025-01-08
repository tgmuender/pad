package user

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"xgmdr.com/pad/internal/api"
	"xgmdr.com/pad/internal/logger"
	pb "xgmdr.com/pad/proto"
)

type Api struct {
	pb.UnimplementedUserServiceServer
}

func (m *Api) WhoAmI(grpcContext context.Context, request *emptypb.Empty) (*pb.UserResponse, error) {
	logger.Get().Debug("User info request received")

	authentication := api.ExtractAuthentication(grpcContext)

	return &pb.UserResponse{Email: authentication.Email}, nil
}
