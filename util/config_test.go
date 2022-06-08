package util

import (
	"context"
	"reflect"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		ctx    context.Context
		logger runtime.Logger
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			name: "NewConfig",
			args: args{
				ctx:    context.Background(),
				logger: &TestLogger{},
			},
			want: &Config{
				InstrumentationName: "nakama-modules-go",
				JaegerURL:           "http://otelcol:14268/api/traces",
				RedpandaURL:         "http://redpanda:8082/topics/nakama",
				ServiceInstanceId:   "00000000-0000-0000-0000-000000000000",
				ServiceName:         "nakama-modules-go",
				ServiceNamespace:    "nakama",
				ServiceVersion:      "v0.1.0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := Config{}
			if appConfig := AppConfig; !reflect.DeepEqual(appConfig, config) {
				t.Errorf("AppConfig = %v, want %v", appConfig, config)
			}
			NewConfig(tt.args.ctx, tt.args.logger)
			if instrumentationName := AppConfig.InstrumentationName; !reflect.DeepEqual(instrumentationName, tt.want.InstrumentationName) {
				t.Errorf("AppConfig = %v, want %v", instrumentationName, tt.want.InstrumentationName)
			}
			if jaegerURL := AppConfig.JaegerURL; !reflect.DeepEqual(jaegerURL, tt.want.JaegerURL) {
				t.Errorf("AppConfig = %v, want %v", jaegerURL, tt.want.JaegerURL)
			}
			if redpandaURL := AppConfig.RedpandaURL; !reflect.DeepEqual(redpandaURL, tt.want.RedpandaURL) {
				t.Errorf("AppConfig = %v, want %v", redpandaURL, tt.want.RedpandaURL)
			}
			if serviceInstanceId := AppConfig.ServiceInstanceId; !reflect.DeepEqual(serviceInstanceId, tt.want.ServiceInstanceId) {
				t.Errorf("AppConfig = %v, want %v", serviceInstanceId, tt.want.ServiceInstanceId)
			}
			if serviceName := AppConfig.ServiceName; !reflect.DeepEqual(serviceName, tt.want.ServiceName) {
				t.Errorf("AppConfig = %v, want %v", serviceName, tt.want.ServiceName)
			}
			if serviceNamespace := AppConfig.ServiceNamespace; !reflect.DeepEqual(serviceNamespace, tt.want.ServiceNamespace) {
				t.Errorf("AppConfig = %v, want %v", serviceNamespace, tt.want.ServiceNamespace)
			}
			if serviceVersion := AppConfig.ServiceVersion; !reflect.DeepEqual(serviceVersion, tt.want.ServiceVersion) {
				t.Errorf("AppConfig = %v, want %v", serviceVersion, tt.want.ServiceVersion)
			}
		})
	}
}
