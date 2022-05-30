package util

import (
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	InstrumentationName string `default:"nakama-modules-go" json:"instrumentation_name,omitempty" required:"true" split_words:"true"`
	JaegerURL           string `default:"http://otelcol:14268/api/traces" json:"jaeger_url,omitempty" required:"true" split_words:"true"`
	RedpandaURL         string `default:"http://redpanda:8082/topics/nakama" json:"redpanda_url,omitempty" required:"true" split_words:"true"`
	ServiceInstanceId   string `default:"00000000-0000-0000-0000-000000000000" json:"service_instance_id,omitempty" required:"true" split_words:"true"`
	ServiceName         string `default:"nakama-modules-go" json:"service_name,omitempty" required:"true" split_words:"true"`
	ServiceNamespace    string `default:"nakama" json:"service_namespace,omitempty" required:"true" split_words:"true"`
	ServiceVersion      string `default:"v0.1.0" json:"service_version,omitempty" required:"true" split_words:"true"`
}

func LoadConfig(logger runtime.Logger) *Config {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		logger.WithField("error", err).Error("Failed to process environment variables")
	}
	return &config
}
