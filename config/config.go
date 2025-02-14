package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config holds the application's environment configurations.
type Config struct {
	Port          int    `envconfig:"API_PORT" default:"80"`
	PodName       string `envconfig:"POD_NAME" default:"QUIZ"`
	RedisKey      string `envconfig:"REDIS_KEY" default:"quiz_leaderboard"`
	RedisAddr     string `envconfig:"REDIS_ADDR" default:"redis:6379"`
	RedisPassword string `envconfig:"REDIS_PASSWORD" default:""`
	RedisDB       int    `envconfig:"REDIS_DB" default:"0"`
}

// LoadConfig retrieves the application's configuration from environment variables.
// It returns a Config struct populated with values from the environment or defaults if not set.
func LoadConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
