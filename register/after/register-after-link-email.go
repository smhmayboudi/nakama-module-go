package after

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func RegisterAfterLinkEmail(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountEmail) error {
	ctx = u.Extract(ctx)
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterLinkEmail",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterLinkEmail", "in": in}); err != nil {
		logger.WithFields(u.InjectMultipleField(ctx)).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
