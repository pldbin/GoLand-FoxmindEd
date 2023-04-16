package logger

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type mockLogger struct {
	mock.Mock
}

func (m *mockLogger) Infof(format string, args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) Errorf(format string, args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) Info(args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) Error(args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) Warnf(format string, args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) Debug(args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) Debugf(format string, args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) WarnfCtx(ctx context.Context, format string, args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) DebugfCtx(ctx context.Context, format string, args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) InfofCtx(ctx context.Context, format string, args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) ErrorfCtx(ctx context.Context, format string, args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) DebugCtx(ctx context.Context, args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) InfoCtx(ctx context.Context, args ...interface{}) {
	// no need any logic yet
}

func (m *mockLogger) ErrorCtx(ctx context.Context, args ...interface{}) {
	// no need any logic yet
}

// NewMock creates logger which does not write logs (useful for tests)
func NewMock() Logger {
	return &mockLogger{}
}
