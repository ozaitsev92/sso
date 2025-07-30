package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/ozaitsev92/sso/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(logger *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.New(logger, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
