package main

import (
	"context"
	"log"
	"log/slog" // <- правильный импорт
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"github.com/1TyuR1/eduquest-backend/services/auth-service/config"
	"github.com/1TyuR1/eduquest-backend/shared/pkg/database"
	"github.com/1TyuR1/eduquest-backend/shared/pkg/logger"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Исправлено: правильный тип
	var appLogger *slog.Logger
	if cfg.Env == "production" {
		appLogger = logger.NewProduction()
	} else {
		appLogger = logger.NewDefault()
	}
	appLogger = logger.WithService(appLogger, "auth-service")

	appLogger.Info("Starting auth-service",
		slog.String("env", cfg.Env),
		slog.String("http_port", cfg.HTTPPort),
		slog.String("grpc_port", cfg.GRPCPort),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbConfig := database.DefaultConfig()
	dbConfig.Host = cfg.DB.Host
	dbConfig.Port = cfg.DB.Port
	dbConfig.User = cfg.DB.User
	dbConfig.Password = cfg.DB.Password
	dbConfig.Database = cfg.DB.Database
	dbConfig.SSLMode = cfg.DB.SSLMode

	dbPool, err := database.NewPostgresPool(ctx, dbConfig, appLogger)
	if err != nil {
		appLogger.Error("Failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer dbPool.Close()

	appLogger.Info("Database connection established")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	appLogger.Info("Shutting down auth-service...")
	cancel()

	appLogger.Info("Auth-service stopped")
}
