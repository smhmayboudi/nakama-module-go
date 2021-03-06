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

func mapper(key string) string {
	switch key {
	case b3ContextHeader:
		return "b3"
	case b3DebugFlagHeader:
		return "flags"
	case b3ParentSpanIDHeader:
		return "parent_span_id"
	case b3SampledHeader:
		return "sampled"
	case b3SpanIDHeader:
		return "span_id"
	case b3TraceIDHeader:
		return "trace_id"
	default:
		return key
	}
}

func Extract(ctx context.Context, encoding b3.Encoding) context.Context {
	// TODO: FIXME
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_VARS).(map[string]string)
	if !ok {
		vars = map[string]string{}
	}
	b3 := b3.New(b3.WithInjectEncoding(encoding))
	return b3.Extract(ctx, propagation.MapCarrier(vars))
}

func Inject(ctx context.Context, encoding b3.Encoding) map[string]interface{} {
	vars := map[string]string{}
	b3 := b3.New(b3.WithInjectEncoding(encoding))
	b3.Inject(ctx, propagation.MapCarrier(vars))
	maps := make(map[string]interface{}, len(vars))
	for k, v := range vars {
		maps[mapper(k)] = v
	}
	return maps
}
