package before

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterBeforeChannelLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *rtapi.Envelope) (*rtapi.Envelope, error) {
	ctx = u.Extract(ctx)
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeChannelLeave",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeChannelLeave", "in": in}); err != nil {
	// 	u.HandleError(ctx, logger, span, err, "Error calling redpanda")
	// 	return in, err
	// }
	return in, nil
}
