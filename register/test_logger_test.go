package register

import "github.com/heroiclabs/nakama-common/runtime"

type TestLogger struct{}

// Debug implements runtime.Logger
func (testLogger *TestLogger) Debug(format string, v ...interface{}) {
}

// Error implements runtime.Logger
func (testLogger *TestLogger) Error(format string, v ...interface{}) {
}

// Fields implements runtime.Logger
func (testLogger *TestLogger) Fields() map[string]interface{} {
	return map[string]interface{}{}
}

// Info implements runtime.Logger
func (testLogger *TestLogger) Info(format string, v ...interface{}) {
}

// Warn implements runtime.Logger
func (testLogger *TestLogger) Warn(format string, v ...interface{}) {
}

// WithField implements runtime.Logger
func (testLogger *TestLogger) WithField(key string, v interface{}) runtime.Logger {
	return testLogger
}

// WithFields implements runtime.Logger
func (testLogger *TestLogger) WithFields(fields map[string]interface{}) runtime.Logger {
	return testLogger
}

var _ runtime.Logger = (*TestLogger)(nil)
