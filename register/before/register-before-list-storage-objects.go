package before

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterBeforeListStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListStorageObjectsRequest) (*api.ListStorageObjectsRequest, error) {
	ctx = u.Extract(ctx, b3.B3SingleHeader)
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListStorageObjects",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListStorageObjects", "in": in}); err != nil {
	// 	u.HandleError(ctx, logger, span, err, "Error calling redpanda")
	// 	return in, err
	// }
	return in, nil
}
