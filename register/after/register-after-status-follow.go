package after

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/nakama-modules-go/util"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func RegisterAfterStatusFollow(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out, in *rtapi.Envelope) error {
	logger.WithFields(u.Inject(ctx, b3.B3MultipleHeader)).WithFields(map[string]interface{}{"name": "RegisterAfterStatusFollow", "in": in, "out": out}).Debug("")
	ctx = u.Extract(ctx, b3.B3SingleHeader)
	ctx, span := otel.Tracer(u.LoadConfig(logger).InstrumentationName).Start(
		ctx,
		"RegisterAfterStatusFollow",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterStatusFollow", "in": in, "out": out}); err != nil {
		logger.WithFields(u.Inject(ctx, b3.B3MultipleHeader)).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
