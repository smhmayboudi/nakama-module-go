package util

import (
	"context"
	"os"
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
		name    string
		args    args
		want    *Config
		wantErr bool
		init    func()
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
				Test:                true,
			},
			wantErr: false,
			init: func() {
				os.Setenv("TEST", "true")
				ModuleConfig = Config{}
			},
		},
		{
			name: "NewConfigWithError",
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
				Test:                true,
			},
			wantErr: true,
			init: func() {
				os.Setenv("TEST", "")
				ModuleConfig = Config{}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init()
			config := Config{}
			if appConfig := ModuleConfig; !reflect.DeepEqual(appConfig, config) {
				t.Errorf("AppConfig = %v, want %v", appConfig, config)
			}
			if err := NewConfig(tt.args.ctx, tt.args.logger); (err != nil) != tt.wantErr {
				t.Errorf("AppConfig error = %v, wantErr %v", err, tt.wantErr)
			}
			if instrumentationName := ModuleConfig.InstrumentationName; !reflect.DeepEqual(instrumentationName, tt.want.InstrumentationName) {
				t.Errorf("AppConfig = %v, want %v", instrumentationName, tt.want.InstrumentationName)
			}
			if jaegerURL := ModuleConfig.JaegerURL; !reflect.DeepEqual(jaegerURL, tt.want.JaegerURL) {
				t.Errorf("AppConfig = %v, want %v", jaegerURL, tt.want.JaegerURL)
			}
			if redpandaURL := ModuleConfig.RedpandaURL; !reflect.DeepEqual(redpandaURL, tt.want.RedpandaURL) {
				t.Errorf("AppConfig = %v, want %v", redpandaURL, tt.want.RedpandaURL)
			}
			if serviceInstanceId := ModuleConfig.ServiceInstanceId; !reflect.DeepEqual(serviceInstanceId, tt.want.ServiceInstanceId) {
				t.Errorf("AppConfig = %v, want %v", serviceInstanceId, tt.want.ServiceInstanceId)
			}
			if serviceName := ModuleConfig.ServiceName; !reflect.DeepEqual(serviceName, tt.want.ServiceName) {
				t.Errorf("AppConfig = %v, want %v", serviceName, tt.want.ServiceName)
			}
			if serviceNamespace := ModuleConfig.ServiceNamespace; !reflect.DeepEqual(serviceNamespace, tt.want.ServiceNamespace) {
				t.Errorf("AppConfig = %v, want %v", serviceNamespace, tt.want.ServiceNamespace)
			}
			if serviceVersion := ModuleConfig.ServiceVersion; !reflect.DeepEqual(serviceVersion, tt.want.ServiceVersion) {
				t.Errorf("AppConfig = %v, want %v", serviceVersion, tt.want.ServiceVersion)
			}
		})
	}
}
