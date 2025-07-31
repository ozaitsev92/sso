package auth

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/ozaitsev92/sso/internal/services/auth"
	"github.com/ozaitsev92/sso/internal/storage"
	"github.com/ozaitsev92/ssoprotos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Register(ctx context.Context, email string, password string) (int64, error)
	Login(ctx context.Context, email string, password string, appID int64) (string, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type serverAPI struct {
	sso.UnimplementedAuthServer
	auth     Auth
	validate *validator.Validate
}

func Register(gRPC *grpc.Server, auth Auth) {
	sso.RegisterAuthServer(gRPC, &serverAPI{
		auth:     auth,
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

	userID, err := s.auth.Register(ctx, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			return nil, status.Errorf(
				codes.AlreadyExists,
				"user already exists: %v",
				err.Error(),
			)
		}

		return nil, status.Errorf(
			codes.Internal,
			"failed to register user: %v",
			err.Error(),
		)
	}

	return &sso.RegisterResponse{
		UserId: userID,
	}, nil
}

type loginInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
	AppId    int64  `validate:"required"`
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

	token, err := s.auth.Login(ctx, input.Email, input.Password, input.AppId)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"invalid credentials: %v",
				err.Error(),
			)
		}

		return nil, status.Errorf(
			codes.Internal,
			"failed to login user: %v",
			err.Error(),
		)
	}

	return &sso.LoginResponse{
		Token: token,
	}, nil
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

	isAdmin, err := s.auth.IsAdmin(ctx, input.UserId)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			return nil, status.Errorf(
				codes.NotFound,
				"user not found: %v",
				err.Error(),
			)
		}

		return nil, status.Errorf(
			codes.Internal,
			"failed to check admin status: %v",
			err.Error(),
		)
	}

	return &sso.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}
