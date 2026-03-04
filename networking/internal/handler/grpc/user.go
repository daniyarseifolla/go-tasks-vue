package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"user-service/internal/model"
	"user-service/internal/service"
	"user-service/gen/pb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) CreateUser(_ context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user, err := h.svc.Create(req.GetName(), req.GetEmail(), req.GetAge())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return toProto(user), nil
}

func (h *UserHandler) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := h.svc.GetByID(req.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return toProto(user), nil
}

func (h *UserHandler) UpdateUser(_ context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	user, err := h.svc.Update(req.GetId(), req.GetName(), req.GetEmail(), req.GetAge())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return toProto(user), nil
}

func (h *UserHandler) DeleteUser(_ context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	if err := h.svc.Delete(req.GetId()); err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &pb.DeleteUserResponse{Success: true}, nil
}

func toProto(u *model.User) *pb.UserResponse {
	return &pb.UserResponse{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Age:       u.Age,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}
