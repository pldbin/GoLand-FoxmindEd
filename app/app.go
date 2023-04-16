package app

import (
	"GoLand-FoxmindEd/config"
	"GoLand-FoxmindEd/logger"
)

func Run() error {
	cfg, err := config.NewConfig()
	if err != nil {
		logger.DebugError("config.NewConfig() failed. Error: ", err)
		return err
	}

	lgr, err := logger.New(&logger.Config{
		LogLevel:    cfg.LogLevel,
		LogServer:   cfg.LogServer,
		ServiceName: cfg.ServiceName,
	})
	if err != nil {
		logger.DebugError("logger.New(), Error: ", err)
		return err
	}
	if err != nil {
		lgr.Infof("Settings loaded, version of Library: v1.0.2")
		return err
	}
	//if err := bot.Add(cfg, *lgr); err != nil {
	//	return err
	//}
	return err
}
