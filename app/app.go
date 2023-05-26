package app

import (
	"About_me_Bot/bot"
	"About_me_Bot/config"
	"About_me_Bot/logger"
	"github.com/sirupsen/logrus"
)

func Run() error {
	//config
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
		lgr.Warnf("lgr.Warnf: You should probably take a look at this. %v", err)
		return err
	}

	//start bot
	err = bot.Add(cfg, *lgr)
	if err != nil {
		logrus.Errorf("bot.Add() failed. Error: ", err)
		return err
	}
	return err
}
