package after

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	u "github.com/smhmayboudi/nakama-modules-go/util"
)

func RegisterAfterValidatePurchaseGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.ValidatePurchaseResponse, in *api.ValidatePurchaseGoogleRequest) error {
	ctx = u.Extract(ctx, b3.B3SingleHeader)
	nakamaContext := u.NewContext(ctx, logger)
	fields := map[string]interface{}{"name": "RegisterAfterValidatePurchaseGoogle", "ctx": nakamaContext, "in": in, "out": out}
	logger.WithFields(u.Inject(ctx, b3.B3MultipleHeader)).WithFields(fields).Debug("")
	ctx, span := otel.Tracer(u.ModuleConfig.InstrumentationName).Start(
		ctx,
		"RegisterAfterValidatePurchaseGoogle",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.RedpandaSend(ctx, logger, fields); err != nil {
		logger.WithFields(u.Inject(ctx, b3.B3MultipleHeader)).WithFields(fields).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return u.InternalError
	}
	return nil
}
