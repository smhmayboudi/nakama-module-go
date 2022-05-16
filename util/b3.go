package util

import (
	"context"

	"github.com/heroiclabs/nakama-common/runtime"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel/propagation"
)

const (
	b3ContextHeader      = "b3"
	b3DebugFlagHeader    = "x-b3-flags"
	b3ParentSpanIDHeader = "x-b3-parentspanid"
	b3SampledHeader      = "x-b3-sampled"
	b3SpanIDHeader       = "x-b3-spanid"
	b3TraceIDHeader      = "x-b3-traceid"
)

func mapper() map[string]string {
	return map[string]string{
		b3ContextHeader:      "b3",
		b3DebugFlagHeader:    "flags",
		b3ParentSpanIDHeader: "parent_span_id",
		b3SampledHeader:      "sampled",
		b3SpanIDHeader:       "span_id",
		b3TraceIDHeader:      "trace_id",
	}
}

func Extract(ctx context.Context) context.Context {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_VARS).(map[string]string)
	if !ok {
		vars = map[string]string{}
	}
	b3 := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader | b3.B3SingleHeader))
	return b3.Extract(ctx, propagation.MapCarrier(vars))
}

func InjectMultipleField(ctx context.Context) map[string]interface{} {
	vars := map[string]string{}
	b3 := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader))
	b3.Inject(ctx, propagation.MapCarrier(vars))
	maps := make(map[string]interface{}, len(vars))
	for k, v := range vars {
		maps[mapper()[k]] = v
	}
	return maps
}

func InjectSingleField(ctx context.Context) map[string]interface{} {
	vars := map[string]string{}
	b3 := b3.New(b3.WithInjectEncoding(b3.B3SingleHeader))
	b3.Fields()
	b3.Inject(ctx, propagation.MapCarrier(vars))
	maps := make(map[string]interface{}, len(vars))
	for k, v := range vars {
		maps[mapper()[k]] = v
	}
	return maps
}
