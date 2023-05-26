package logger

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

const requestID = "request_id"

// getID will return request_id from context
func getID(ctx context.Context) string {
	reqID, ok := ctx.Value(requestID).(string)
	if !ok {
		return ""
	}
	return reqID
}

// funcName will send to logger name of func, which logger will print
func funcName(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	nameEnd := filepath.Ext(runtime.FuncForPC(pc).Name())
	return fmt.Sprintf("Func: %v()", strings.TrimPrefix(nameEnd, "."))
}

// Infof writes debug info with given format
func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	l.entry.Infof(fmt.Sprintf("%s %v", funcName(2), format), args...)
}

// Errorf writes debug error with given format
func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(fmt.Sprintf("%s %v", funcName(2), format), args...)
}

// Info writes debug info
func (l *LogrusLogger) Info(args ...interface{}) {
	l.entry.Info(append([]interface{}{funcName(2)}, args...)...)
}

// Error writes debug error
func (l *LogrusLogger) Error(args ...interface{}) {
	l.entry.Error(append([]interface{}{funcName(2)}, args...)...)
}

// Warnf writes debug message
func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

// Debug writes debug message
func (l *LogrusLogger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

// Debugf writes debug message
func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

// WarnfCtx writes warning message with given format and 'request_id' from context.
func (l *LogrusLogger) WarnfCtx(ctx context.Context, format string, args ...interface{}) {
	id := getID(ctx)
	if id == "" {
		l.entry.Warnf(fmt.Sprintf("%s %v", funcName(2), format), args...)
		return
	}
	l.entry.WithField(requestID, id).
		Warnf(fmt.Sprintf("%s %v", funcName(2), format), args...)
}

// DebugfCtx writes debug message with given format and 'request_id' from context.
func (l *LogrusLogger) DebugfCtx(ctx context.Context, format string, args ...interface{}) {
	id := getID(ctx)
	if id == "" {
		l.entry.Debugf(fmt.Sprintf("%s %v", funcName(2), format), args...)
		return
	}
	l.entry.WithField(requestID, id).
		Debugf(fmt.Sprintf("%s %v", funcName(2), format), args...)
}

// InfofCtx writes debug info with given format and 'request_id' from context.
func (l *LogrusLogger) InfofCtx(ctx context.Context, format string, args ...interface{}) {
	id := getID(ctx)
	if id == "" {
		l.entry.Infof(fmt.Sprintf("%s %v", funcName(2), format), args...)
		return
	}
	l.entry.WithField(requestID, id).
		Infof(fmt.Sprintf("%s %v", funcName(2), format), args...)
}

// ErrorfCtx writes debug error with given format and 'request_id' from context.
func (l *LogrusLogger) ErrorfCtx(ctx context.Context, format string, args ...interface{}) {
	id := getID(ctx)
	if id == "" {
		l.entry.Errorf(fmt.Sprintf("%s %v", funcName(2), format), args...)
		return
	}
	l.entry.WithField(requestID, id).
		Errorf(fmt.Sprintf("%s %v", funcName(2), format), args...)
}

// DebugCtx writes debug message with 'request_id' from context.
func (l *LogrusLogger) DebugCtx(ctx context.Context, args ...interface{}) {
	id := getID(ctx)
	if id == "" {
		l.entry.Debug(append([]interface{}{funcName(2)}, args...)...)
		return
	}
	l.entry.WithField(requestID, id).
		Debug(append([]interface{}{funcName(2)}, args...)...)
}

// InfoCtx writes debug info with 'request_id' from context.
func (l *LogrusLogger) InfoCtx(ctx context.Context, args ...interface{}) {
	id := getID(ctx)
	if id == "" {
		l.entry.Info(append([]interface{}{funcName(2)}, args...)...)
		return
	}
	l.entry.WithField(requestID, id).
		Info(append([]interface{}{funcName(2)}, args...)...)
}

// ErrorCtx writes debug error with 'request_id' from context.
func (l *LogrusLogger) ErrorCtx(ctx context.Context, args ...interface{}) {
	id := getID(ctx)
	if id == "" {
		l.entry.Error(append([]interface{}{funcName(2)}, args...)...)
		return
	}
	l.entry.WithField(requestID, id).
		Error(append([]interface{}{funcName(2)}, args...)...)
}
