package util

import (
	"context"

	"github.com/heroiclabs/nakama-common/runtime"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel/trace"
)

func NewSpanContext(ctx context.Context, options ...trace.TracerOption) trace.SpanContext {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_VARS).(map[string]string)
	if !ok {
		vars = map[string]string{}
	}
	textMapCarrier := TextMapCarrier{
		B3:           vars["b3"],
		ParentSpanId: vars["x-b3-parentspanid"],
		Sampled:      vars["x-b3-sampled"],
		SpanId:       vars["x-b3-spanid"],
		TraceId:      vars["x-b3-traceid"],
	}
	b3 := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader | b3.B3SingleHeader))
	b3.Extract(ctx, &textMapCarrier)
	spanId, _ := trace.SpanIDFromHex(textMapCarrier.SpanId)
	traceId, _ := trace.TraceIDFromHex(textMapCarrier.TraceId)
	return trace.NewSpanContext(trace.SpanContextConfig{
		SpanID:  spanId,
		TraceID: traceId,
	})
}
