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

func RegisterEventSessionStart(ctx context.Context, logger runtime.Logger, evt *api.Event) {
	ctx = u.Extract(ctx)
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterEventSessionStart",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterEventSessionStart", "event": evt}); err != nil {
		logger.WithFields(u.InjectMultipleField(ctx)).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
	}
	logger.WithFields(u.InjectMultipleField(ctx)).Info(fmt.Sprintf("session start %v %v", ctx, evt))
}
