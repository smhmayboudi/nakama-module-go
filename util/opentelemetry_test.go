package util

import (
	"context"
	"reflect"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func TestNewOpenTelemetry(t *testing.T) {
	type args struct {
		ctx    context.Context
		logger runtime.Logger
	}
	tests := []struct {
		name string
		args args
		want trace.TracerProvider
	}{
		{
			name: "NewOpenTelemetry",
			args: args{
				ctx:    context.Background(),
				logger: &TestLogger{},
			},
			want: otel.GetTracerProvider(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tp := otel.GetTracerProvider(); !reflect.DeepEqual(tp, tt.want) {
				t.Errorf("NewOpenTelemetry() = %v, want %v", tp, tt.want)
			}
			NewOpenTelemetry(tt.args.ctx, tt.args.logger)
			if tp := otel.GetTracerProvider(); reflect.DeepEqual(tp, tt.want) {
				t.Errorf("NewOpenTelemetry() = %v, want !%v", tp, tt.want)
			}
		})
	}
}
