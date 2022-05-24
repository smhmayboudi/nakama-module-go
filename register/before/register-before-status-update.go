package before

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/nakama-modules-go/util"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterBeforeStatusUpdate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *rtapi.Envelope) (*rtapi.Envelope, error) {
	ctx = u.Extract(ctx, b3.B3SingleHeader)
	_, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeStatusUpdate",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeStatusUpdate", "in": in}); err != nil {
	// 	u.HandleError(ctx, logger, span, err, "Error calling redpanda")
	// 	return in, err
	// }
	return in, nil
}
