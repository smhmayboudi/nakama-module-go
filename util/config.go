package util

import (
	"context"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	InstrumentationName string `default:"nakama-modules-go" envconfig:"INSTRUMENTATION_NAME" json:"instrumentation_name,omitempty" protobuf:"bytes,1,opt,name=instrumentation_name,json=instrumentation_name,proto3" required:"true" split_words:"true"`
	JaegerURL           string `default:"http://otelcol:14268/api/traces" envconfig:"JAEGER_URL" json:"jaeger_url,omitempty" protobuf:"bytes,2,opt,name=jaeger_url,json=jaeger_url,proto3" required:"true" split_words:"true"`
	RedpandaURL         string `default:"http://redpanda:8082/topics/nakama" envconfig:"REDPANDA_URL" json:"redpanda_url,omitempty" protobuf:"bytes,3,opt,name=redpanda_url,json=redpanda_url,proto3" required:"true" split_words:"true"`
	ServiceInstanceId   string `default:"00000000-0000-0000-0000-000000000000" envconfig:"SERVICE_INSTANCE_ID" json:"service_instance_id,omitempty" protobuf:"bytes,4,opt,name=service_instance_id,json=service_instance_id,proto3" required:"true" split_words:"true"`
	ServiceName         string `default:"nakama-modules-go" envconfig:"SERVICE_NAME" json:"service_name,omitempty" protobuf:"bytes,5,opt,name=service_name,json=service_name,proto3" required:"true" split_words:"true"`
	ServiceNamespace    string `default:"nakama" envconfig:"SERVICE_NAMESPACE" json:"service_namespace,omitempty" protobuf:"bytes,6,opt,name=service_namespace,json=service_namespace,proto3" required:"true" split_words:"true"`
	ServiceVersion      string `default:"v0.1.0" envconfig:"SERVICE_VERSION" json:"service_version,omitempty" protobuf:"bytes,7,opt,name=service_version,json=service_version,proto3" required:"true" split_words:"true"`
}

var ModuleConfig Config

func NewConfig(ctx context.Context, logger runtime.Logger) {
	nakamaContext := NewContext(ctx, logger)
	fields := map[string]interface{}{"name": "NewConfig", "ctx": nakamaContext}
	if err := envconfig.Process("", &ModuleConfig); err != nil {
		logger.WithFields(fields).WithField("error", err).Error("Failed to process environment variables")
	}
}
