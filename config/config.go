package config

import (
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	Home         string        `env:"HOME"  envDefault:""`
	Port         int           `env:"PORT" envDefault:"8085"`
	Password     string        `env:"PASSWORD"`
	IsProduction bool          `env:"PRODUCTION"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION"`
	TempFolder   string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
	Btoken       string        `env:"BTOKEN" envDefault:"5844982587:AAFG3r6a7SIi11mrjFYqA_QbgrMacqdPayI"`

	ServiceName string `env:"SERVICE_NAME"`
	LogServer   string `env:"LOG_SERVER"`
	LogLevel    string `env:"LOG_LEVEL"`
}

// NewConfig parses envs and constructs the config
func NewConfig() (*Config, error) {
	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
