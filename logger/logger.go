package logger

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	graylog "gopkg.in/gemnasium/logrus-graylog-hook.v2"
)

const timeLayout = "2006-01-02 15:04:05.000000"

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})

	WarnfCtx(ctx context.Context, format string, args ...interface{})
	DebugfCtx(ctx context.Context, format string, args ...interface{})
	InfofCtx(ctx context.Context, format string, args ...interface{})
	ErrorfCtx(ctx context.Context, format string, args ...interface{})

	DebugCtx(ctx context.Context, args ...interface{})
	InfoCtx(ctx context.Context, args ...interface{})
	ErrorCtx(ctx context.Context, args ...interface{})
}

// Config represents config object for logger package
type Config struct {
	LogLevel    string
	LogServer   string
	ServiceName string
}

// LogrusLogger logger implementation with logrus internals
type LogrusLogger struct {
	logrus *logrus.Logger
	entry  *logrus.Entry
}

// New creates new logger
func New(cfg *Config) (*LogrusLogger, error) {
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	logger := &LogrusLogger{
		logrus: logrus.New(),
	}

	logger.logrus.SetLevel(level)

	customFormatter := &logrus.TextFormatter{
		TimestampFormat: timeLayout,
		ForceColors:     true,
		FullTimestamp:   true,
	}

	logger.logrus.SetFormatter(customFormatter)

	if cfg.LogServer != "" {
		logger.logrus.AddHook(
			graylog.NewGraylogHook(cfg.LogServer, map[string]interface{}{}),
		)
	}

	logger.entry = logger.logrus.WithFields(logrus.Fields{
		"service_name": cfg.ServiceName,
	})

	return logger, nil
}
