package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
	}
	Grinex struct {
		URL string
	}
	Postgres struct {
		DSN string
	}
}

func Load() (*Config, error) {
	var configFile string
	flag.StringVar(&configFile, "config", "config.yaml", "Path to config file")
	flag.Parse()

	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	// ENV overrides
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP") // e.g. APP_GRINEX_URL

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	// Defaults
	viper.SetDefault("server.port", "50051")
	viper.SetDefault("grinex.url", "https://grinex.io/api/v2/depth")
	viper.SetDefault("postgres.dsn", "postgres://postgres:postgres@localhost:5432/usdt_rates?sslmode=disable")

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return &cfg, nil
}
