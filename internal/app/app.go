package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/ozaitsev92/sso/internal/app/grpc"
	"github.com/ozaitsev92/sso/internal/services/auth"
	"github.com/ozaitsev92/sso/internal/storage/sqlite"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(logger *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(logger, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(logger, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
