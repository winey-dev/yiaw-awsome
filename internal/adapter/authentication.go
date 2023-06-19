package adapter

import (
	"context"
	"yiaw-awsome/internal/domain/authentication"
	v1authenticationpb "yiaw-awsome/proto/v1/authentication"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Authentication struct {
	v1authenticationpb.AuthenticationServiceServer
	authenticationService authentication.AuthenticationService
}

func NewAuthentication(authenticationService authentication.AuthenticationService) *Authentication {
	return &Authentication{authenticationService: authenticationService}
}

func (a *Authentication) Login(ctx context.Context, req *v1authenticationpb.LoginRequest) (*v1authenticationpb.LoginResponse, error) {
	tok, err := a.authenticationService.Login(ctx, req.UserId, req.Password)
	if err != nil {
		return &v1authenticationpb.LoginResponse{}, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	return &v1authenticationpb.LoginResponse{
		Bearer: tok,
	}, nil
}

func (a *Authentication) Logout(ctx context.Context, req *v1authenticationpb.LogoutRequest) (*emptypb.Empty, error) {
	if err := a.authenticationService.Logout(ctx, req.Bearer); err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}
	return &emptypb.Empty{}, nil
}

func (a *Authentication) ChangePassword(ctx context.Context, req *v1authenticationpb.ChangePasswordRequest) (*v1authenticationpb.ChangePasswordResponse, error) {
	tok, err := a.authenticationService.ChangePassword(ctx, req.UserId, req.CurrentPassword, req.NewPassword)
	if err != nil {
		return &v1authenticationpb.ChangePasswordResponse{}, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return &v1authenticationpb.ChangePasswordResponse{
		NewBearer: tok,
	}, nil
}
