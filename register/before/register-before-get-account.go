package before

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/nakama-modules-go/util"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterBeforeGetAccount(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) error {
	ctx = u.Extract(ctx, b3.B3SingleHeader)
	_, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeGetAccount",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeGetAccount"}); err != nil {
	// 	u.HandleError(ctx, logger, span, err, "Error calling redpanda")
	// 	return err
	// }
	return nil
}
