package before

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterBeforeLinkFacebook(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LinkFacebookRequest) (*api.LinkFacebookRequest, error) {
	ctx = u.Extract(ctx)
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeLinkFacebook",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeLinkFacebook", "in": in}); err != nil {
	// 	u.HandleError(ctx, logger, span, err, "Error calling redpanda")
	// 	return in, err
	// }
	return in, nil
}
