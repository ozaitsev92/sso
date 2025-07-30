package auth

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/ozaitsev92/ssoprotos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	sso.UnimplementedAuthServer
	validate *validator.Validate
}

func Register(gRPC *grpc.Server) {
	sso.RegisterAuthServer(gRPC, &serverAPI{
		validate: validator.New(),
	})
}

type registerInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func (s *serverAPI) Register(ctx context.Context, req *sso.RegisterRequest) (*sso.RegisterResponse, error) {
	input := registerInput{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	if err := s.validate.Struct(input); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"validation failed: %v",
			err.Error(),
		)
	}

	return &sso.RegisterResponse{}, nil // TODO: Implement login logic
}

type loginInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
	AppId    int32  `validate:"required"`
}

func (s *serverAPI) Login(ctx context.Context, req *sso.LoginRequest) (*sso.LoginResponse, error) {
	input := loginInput{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		AppId:    req.GetAppId(),
	}

	if err := s.validate.Struct(input); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"validation failed: %v",
			err.Error(),
		)
	}

	return &sso.LoginResponse{}, nil // TODO: Implement login logic
}

type isAdminInput struct {
	UserId int64 `validate:"required"`
}

func (s *serverAPI) IsAsmin(ctx context.Context, req *sso.IsAdminRequest) (*sso.IsAdminResponse, error) {
	input := isAdminInput{
		UserId: req.GetUserId(),
	}

	if err := s.validate.Struct(input); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"validation failed: %v",
			err.Error(),
		)
	}

	return &sso.IsAdminResponse{IsAdmin: true}, nil // TODO: Implement admin check logic
}
