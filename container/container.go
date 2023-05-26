package container

import (
	"About_me_Bot/config"
	"About_me_Bot/logger"
)

type BotContainerInterface interface {
	GetConfig() *config.Config
	GetLogger() *logger.Logger
}

// container struct
type container struct {
	config *config.Config
	logger *logger.Logger
}

func NewBotInfrastructureContainer(config *config.Config, logger *logger.Logger) BotContainerInterface {
	return &container{
		config: config,
		logger: logger,
	}
}

// GetConfig returns the object of configuration
func (c *container) GetConfig() *config.Config {
	return c.config
}

// GetLogger returns the object of logger
func (c *container) GetLogger() *logger.Logger {
	return c.logger
}
