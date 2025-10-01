package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type GRPCConfig struct {
	Port    int           `yaml:"port"    env:"GRPC_PORT"    env-default:"2441"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_TIMEOUT" env-default:"5s"`
}

type Config struct {
	Env            string        `yaml:"env"           env:"APP_ENV"        env-default:"local"`
	StoragePath    string        `yaml:"storage_path"  env:"STORAGE_PATH"   env-default:"./storage/sso.db"`
	PostgresDSN    string        `yaml:"postgres_dsn"  env:"POSTGRES_DSN"   env-default:"postgresql://postgres:postgres@localhost:5432/sso?sslmode=verify-ca&sslrootcert=./certs/ca.pem"`
	TokenTTL       time.Duration `yaml:"token_ttl"     env:"TOKEN_TTL"      env-default:"1h"`
	GRPC           GRPCConfig    `yaml:"grpc"`
	MigrationsPath string        `yaml:"-"             env:"MIGRATIONS_PATH" env-default:"./migrations"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}
	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	// Сначала читаем файл, затем переменные окружения перекрывают значения
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}
	// Явно перечитать ENV ещё раз (опционально, на случай отсутствия файла)
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic("cannot read env: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	var showEnv bool
	flag.BoolVar(&showEnv, "env", false, "print env help and exit")
	flag.Parse()

	if showEnv {
		var cfg Config
		text, err := cleanenv.GetDescription(&cfg, nil)
		if err != nil {
			fmt.Println("failed to build env help:", err)
			os.Exit(2)
		}
		fmt.Println("Environment variables:")
		fmt.Println(text)
		os.Exit(0)
	}

	if res == "" {
		res = "./config/config.yaml"
	}
	return res
}
