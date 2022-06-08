package util

import (
	"context"
	"reflect"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel/propagation"
)

func Test_mapper(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "b3ContextHeader",
			args: args{key: b3ContextHeader},
			want: "b3",
		},
		{
			name: "b3DebugFlagHeader",
			args: args{key: b3DebugFlagHeader},
			want: "flags",
		},
		{
			name: "b3ParentSpanIDHeader",
			args: args{key: b3ParentSpanIDHeader},
			want: "parent_span_id",
		},
		{
			name: "b3SampledHeader",
			args: args{key: b3SampledHeader},
			want: "sampled",
		},
		{
			name: "b3SpanIDHeader",
			args: args{key: b3SpanIDHeader},
			want: "span_id",
		},
		{
			name: "b3TraceIDHeader",
			args: args{key: b3TraceIDHeader},
			want: "trace_id",
		},
		{
			name: "default",
			args: args{key: "default"},
			want: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapper(tt.args.key); got != tt.want {
				t.Errorf("mapper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtract(t *testing.T) {
	type args struct {
		ctx      context.Context
		encoding b3.Encoding
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "default",
			args: args{
				ctx:      context.Background(),
				encoding: b3.B3SingleHeader,
			},
			want: func() context.Context {
				ctx := context.Background()
				vars, ok := ctx.Value(runtime.RUNTIME_CTX_VARS).(map[string]string)
				if !ok {
					vars = map[string]string{}
				}
				b3 := b3.New(b3.WithInjectEncoding(b3.B3SingleHeader))
				return b3.Extract(ctx, propagation.MapCarrier(vars))
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Extract(tt.args.ctx, tt.args.encoding); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInject(t *testing.T) {
	type args struct {
		ctx      context.Context
		encoding b3.Encoding
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				ctx:      context.Background(),
				encoding: b3.B3SingleHeader,
			},
			want: func() map[string]interface{} {
				ctx := context.Background()
				vars := map[string]string{}
				b3 := b3.New(b3.WithInjectEncoding(b3.B3SingleHeader))
				b3.Inject(ctx, propagation.MapCarrier(vars))
				maps := make(map[string]interface{}, len(vars))
				for k, v := range vars {
					maps[mapper(k)] = v
				}
				return maps
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Inject(tt.args.ctx, tt.args.encoding); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inject() = %v, want %v", got, tt.want)
			}
		})
	}
}
