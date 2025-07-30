package auth

import (
	"context"

	"github.com/ozaitsev92/ssoprotos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	sso.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	sso.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Register(context.Context, *sso.RegisterRequest) (*sso.RegisterResponse, error) {
	return nil, nil // TODO: Implement registration logic
}

func (s *serverAPI) Login(context.Context, *sso.LoginRequest) (*sso.LoginResponse, error) {
	return nil, nil // TODO: Implement login logic
}

func (s *serverAPI) IsAsmin(context.Context, *sso.IsAdminRequest) (*sso.IsAdminResponse, error) {
	return nil, nil // TODO: Implement admin check logic
}
