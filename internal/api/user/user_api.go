package user

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"xgmdr.com/pad/internal/logger"
	"xgmdr.com/pad/internal/storage"
	pb "xgmdr.com/pad/proto"
)

type Api struct {
	pb.UnimplementedUserServiceServer
}

func (m *Api) WhoAmI(grpcContext context.Context, request *emptypb.Empty) (*pb.UserResponse, error) {
	logger.Get().Debug("User info request received")

	user, ok := grpcContext.Value("user").(*storage.User)
	if !ok {
		logger.Get().Debug("User not found")
		return nil, errors.New("user not found in context")
	}

	return &pb.UserResponse{Email: user.Email}, nil
}
