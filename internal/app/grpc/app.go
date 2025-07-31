package grpcapp

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	authgrpc "github.com/ozaitsev92/sso/internal/grpc/auth"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

type Auth interface {
	Register(ctx context.Context, email string, password string) (int64, error)
	Login(ctx context.Context, email string, password string, appID int64) (string, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

func New(logger *slog.Logger, authService Auth, port int) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer, authService)

	return &App{
		log:        logger,
		port:       port,
		gRPCServer: gRPCServer,
	}
}

func (a *App) Run() error {
	const op = "grpcapp.App.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server stopped")

	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	const op = "grpcapp.App.Stop"

	log := a.log.With(
		slog.String("op", op),
	)

	log.Info("stopping grpc server")

	a.gRPCServer.GracefulStop()

	log.Info("grpc server stopped")
}
