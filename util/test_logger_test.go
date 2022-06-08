package util

import "github.com/heroiclabs/nakama-common/runtime"

type TestLogger struct{}

var testLogger = &TestLogger{}

// Debug implements runtime.Logger
func (*TestLogger) Debug(format string, v ...interface{}) {
}

// Error implements runtime.Logger
func (*TestLogger) Error(format string, v ...interface{}) {
}

// Fields implements runtime.Logger
func (*TestLogger) Fields() map[string]interface{} {
	return map[string]interface{}{}
}

// Info implements runtime.Logger
func (*TestLogger) Info(format string, v ...interface{}) {
}

// Warn implements runtime.Logger
func (*TestLogger) Warn(format string, v ...interface{}) {
}

// WithField implements runtime.Logger
func (*TestLogger) WithField(key string, v interface{}) runtime.Logger {
	return testLogger
}

// WithFields implements runtime.Logger
func (*TestLogger) WithFields(fields map[string]interface{}) runtime.Logger {
	return testLogger
}

var _ runtime.Logger = (*TestLogger)(nil)
