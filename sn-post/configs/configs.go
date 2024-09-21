package configs

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"path"
)

type Config struct {
	App
	HTTP
	Log
	PG
	JWT
}

type App struct {
	Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
	Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
}

type HTTP struct {
	Host string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
	Port string `env_required:"true" yaml:"port" env:"HTTP_PORT"`
}

type Log struct {
	Level string `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
}

type PG struct {
	MaxPoolSize int    `env-required:"true" yaml:"max_pool_size" env:"PG_MAX_POOL_SIZE"`
	URL         string `env-required:"true" yaml:"url" env:"PG_URL"`
}

type JWT struct {
	SignKey  string `env-required:"true" yaml:"sign_key" env:"JWT_SIGN_KEY"`
	TokenTTL string `env-required:"true" yaml:"token_ttl" env:"JWT_TOKEN_TTL"`
}

func New(pathConfig string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path.Join("./", pathConfig), cfg)
	if err != nil {
		return nil, fmt.Errorf("error read configs: %v", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("Error update env: %v", err)
	}

	return cfg, nil
}
