package util

import "github.com/heroiclabs/nakama-common/runtime"

type TestLogger struct{}

// Debug implements runtime.Logger
func (l *TestLogger) Debug(format string, v ...interface{}) {
}

// Error implements runtime.Logger
func (l *TestLogger) Error(format string, v ...interface{}) {
}

// Fields implements runtime.Logger
func (l *TestLogger) Fields() map[string]interface{} {
	return map[string]interface{}{}
}

// Info implements runtime.Logger
func (l *TestLogger) Info(format string, v ...interface{}) {
}

// Warn implements runtime.Logger
func (l *TestLogger) Warn(format string, v ...interface{}) {
}

// WithField implements runtime.Logger
func (l *TestLogger) WithField(key string, v interface{}) runtime.Logger {
	return l
}

// WithFields implements runtime.Logger
func (l *TestLogger) WithFields(fields map[string]interface{}) runtime.Logger {
	return l
}

var _ runtime.Logger = (*TestLogger)(nil)
