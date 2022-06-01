package util

import (
	"context"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

func NewOpenTelemetry(ctx context.Context, logger runtime.Logger) func() {
	nakamaContext := NewContext(ctx, logger)
	fields := map[string]interface{}{"name": "NewOpenTelemetry", "ctx": nakamaContext}
	rna := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceInstanceIDKey.String(AppConfig.ServiceInstanceId),
		semconv.ServiceNameKey.String(AppConfig.ServiceName),
		semconv.ServiceNamespaceKey.String(AppConfig.ServiceNamespace),
		semconv.ServiceVersionKey.String(AppConfig.ServiceVersion),
	)
	rm, err := resource.Merge(
		resource.Default(),
		rna,
	)
	if err != nil {
		logger.WithFields(fields).WithField("error", err).Error("Failed to merge resources")
	}
	rn, err := resource.New(
		ctx,
		// resource.WithAttributes(),
		resource.WithContainer(),
		// resource.WithDetectors(thirdparty.Detector{}),
		resource.WithFromEnv(),
		resource.WithOS(),
		resource.WithProcess(),
	)
	if err != nil {
		logger.WithFields(fields).WithField("error", err).Error("Failed to create a resource")
	}
	r, err := resource.Merge(
		rm,
		rn,
	)
	if err != nil {
		logger.WithFields(fields).WithField("error", err).Error("Failed to merge resources")
	}
	ej, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(AppConfig.JaegerURL)))
	if err != nil {
		logger.WithFields(fields).WithField("error", err).Error("Failed to create jaeger exporter")
	}
	es, err := stdouttrace.New()
	if err != nil {
		logger.WithFields(fields).WithField("error", err).Error("Failed to create stdouttrace exporter")
	}
	tp := sdkTrace.NewTracerProvider(
		sdkTrace.WithBatcher(ej),
		sdkTrace.WithResource(r),
		sdkTrace.WithSyncer(es),
	)
	otel.SetTracerProvider(tp)

	return func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		defer func(ctx context.Context) {
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			if err := tp.Shutdown(ctx); err != nil {
				logger.WithFields(fields).WithField("error", err).Error("Failed to shutdown tracer provider")
			}
		}(ctx)
	}
}
