package grpc

import (
	"context"
	userPB "movie/user/internal/pb/user"
	"movie/user/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserGRPCHandler struct {
	userPB.UnimplementedUserServer
	service *service.Service
}

func NewUserGRPCHandler(service *service.Service) *UserGRPCHandler {
	return &UserGRPCHandler{
		service: service,
	}
}

func (h *UserGRPCHandler) Signup(context.Context, *userPB.NewUserInput) (*userPB.APICommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}

func (h *UserGRPCHandler) Login(context.Context, *userPB.LoginInput) (*userPB.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
