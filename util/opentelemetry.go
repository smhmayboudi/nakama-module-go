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
		semconv.ServiceInstanceIDKey.String(ModuleConfig.ServiceInstanceId),
		semconv.ServiceNameKey.String(ModuleConfig.ServiceName),
		semconv.ServiceNamespaceKey.String(ModuleConfig.ServiceNamespace),
		semconv.ServiceVersionKey.String(ModuleConfig.ServiceVersion),
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
	ej, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(ModuleConfig.JaegerURL)))
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
		ctxWC, cancel := context.WithCancel(ctx)
		defer cancel()
		defer func(ctxWT context.Context) {
			ctxWT, cancel := context.WithTimeout(ctxWT, 5*time.Second)
			defer cancel()
			if err := tp.Shutdown(ctxWT); err != nil {
				logger.WithFields(fields).WithField("error", err).Error("Failed to shutdown tracer provider")
			}
		}(ctxWC)
	}
}
