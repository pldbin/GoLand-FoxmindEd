package config

import (
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	Home         string        `env:"HOME"  envDefault:""`
	Port         int           `env:"PORT" envDefault:"8085"`
	Password     string        `env:"PASSWORD" envDefault:""`
	IsProduction bool          `env:"PRODUCTION" envDefault:"false"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION" envDefault:"5m"`
	TempFolder   string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
	Btoken       string        `env:"BTOKEN" envDefault:"5844982587:AAFG3r6a7SIi11mrjFYqA_QbgrMacqdPayI"`
	HolidaysAPI  string        `env:"HOLIDAY" envDefault:"1a3e1df232c84346b2b2d11d78ce89f2"`

	ServiceName string `env:"SERVICE_NAME"envDefault:"about-me-bot"`
	LogServer   string `env:"LOG_SERVER" envDefault:"stdout"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
}

// NewConfig parses envs and constructs the config
func NewConfig() (*Config, error) {
	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
