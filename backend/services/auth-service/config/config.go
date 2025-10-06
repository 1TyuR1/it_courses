// services/auth-service/config/config.go
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Env string

	// Server
	HTTPPort string
	GRPCPort string

	// Database
	DB DBConfig

	// Redis
	Redis RedisConfig

	// JWT
	JWT JWTConfig

	// OAuth
	Google   OAuthConfig
	Apple    OAuthConfig
	Telegram TelegramConfig
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type JWTConfig struct {
	Secret     string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

type TelegramConfig struct {
	BotToken string
}

// Load загружает конфигурацию из environment variables
func MustLoad() (*Config, error) {
	cfg := &Config{
		Env: getEnv("ENV", "development"),

		HTTPPort: getEnv("AUTH_SERVICE_HTTP_PORT", "8081"),
		GRPCPort: getEnv("AUTH_SERVICE_GRPC_PORT", "9081"),

		DB: DBConfig{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnvAsInt("POSTGRES_PORT", 5432),
			User:     getEnv("POSTGRES_USER", "eduquest"),
			Password: getEnv("POSTGRES_PASSWORD", "dev_password"),
			Database: getEnv("AUTH_DB_NAME", "auth_db"),
			SSLMode:  getEnv("POSTGRES_SSL_MODE", "disable"),
		},

		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},

		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", ""),
			AccessTTL:  getEnvAsDuration("JWT_ACCESS_TTL", "15m"),
			RefreshTTL: getEnvAsDuration("JWT_REFRESH_TTL", "720h"),
		},

		Google: OAuthConfig{
			ClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
			ClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
			RedirectURL:  getEnv("GOOGLE_REDIRECT_URL", "http://localhost:8081/auth/google/callback"),
		},

		Apple: OAuthConfig{
			ClientID:     getEnv("APPLE_CLIENT_ID", ""),
			ClientSecret: getEnv("APPLE_CLIENT_SECRET", ""),
			RedirectURL:  getEnv("APPLE_REDIRECT_URL", "http://localhost:8081/auth/apple/callback"),
		},

		Telegram: TelegramConfig{
			BotToken: getEnv("TELEGRAM_BOT_TOKEN", ""),
		},
	}

	// Валидация обязательных полей
	if cfg.JWT.Secret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	return cfg, nil
}

// Helper функции
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue string) time.Duration {
	valueStr := getEnv(key, defaultValue)
	duration, err := time.ParseDuration(valueStr)
	if err != nil {
		duration, _ = time.ParseDuration(defaultValue)
	}
	return duration
}
