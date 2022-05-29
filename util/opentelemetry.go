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

func InitProvider(ctx context.Context, logger runtime.Logger) func() {
	rna := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceInstanceIDKey.String(LoadConfig(logger).ServiceInstanceId),
		semconv.ServiceNameKey.String(LoadConfig(logger).ServiceName),
		semconv.ServiceNamespaceKey.String(LoadConfig(logger).ServiceNamespace),
		semconv.ServiceVersionKey.String(LoadConfig(logger).ServiceVersion),
	)
	rm, err := resource.Merge(
		resource.Default(),
		rna,
	)
	if err != nil {
		logger.WithField("error", err).Error("Failed to merge resources")
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
		logger.WithField("error", err).Error("Failed to create a resource")
	}
	r, err := resource.Merge(
		rm,
		rn,
	)
	if err != nil {
		logger.WithField("error", err).Error("Failed to merge resources")
	}
	ej, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(LoadConfig(logger).JaegerURL)))
	if err != nil {
		logger.WithField("error", err).Error("Failed to create jaeger exporter")
	}
	es, err := stdouttrace.New()
	if err != nil {
		logger.WithField("error", err).Error("Failed to create stdouttrace exporter")
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
				logger.WithField("error", err).Error("Failed to shutdown tracer provider")
			}
		}(ctx)
	}
}
