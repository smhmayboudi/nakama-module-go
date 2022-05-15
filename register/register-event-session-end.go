package register

import (
	"context"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func RegisterEventSessionEnd(ctx context.Context, logger runtime.Logger, evt *api.Event) {
	ctx = trace.ContextWithRemoteSpanContext(ctx, u.NewSpanContext(ctx))
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterEventSessionEnd",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterEventSessionEnd", "event": evt}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
	}
	textMapCarrier := u.NewTextMapCarrier(ctx)
	logger.WithFields(textMapCarrier.MultipleField()).Info(fmt.Sprintf("session end %v %v", ctx, evt))
}
