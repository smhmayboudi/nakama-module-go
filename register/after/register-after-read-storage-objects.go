package after

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func RegisterAfterReadStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.StorageObjects, in *api.ReadStorageObjectsRequest) error {
	ctx = u.Extract(ctx, b3.B3SingleHeader)
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterReadStorageObjects",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		logger.WithFields(u.Inject(ctx, b3.B3MultipleHeader)).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
