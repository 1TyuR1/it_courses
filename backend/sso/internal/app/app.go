package app

import (
	"context"
	"log/slog"
	"time"

	grpcapp "sso/internal/app/grpc"
	"sso/internal/service/auth"
	"sso/internal/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	connString string,
	tokenTTL time.Duration,
) *App {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	storage, err := postgres.New(ctx, connString)
	if err != nil {
		// в проде лучше вернуть error наружу, здесь оставим panic как было
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
