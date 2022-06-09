package register

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

func RegisterTournamentReset(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, tournament *api.Tournament, end, reset int64) error {
	ctx = u.Extract(ctx, b3.B3SingleHeader)
	nakamaContext := u.NewContext(ctx, logger)
	fields := map[string]interface{}{"name": "RegisterTournamentReset", "ctx": nakamaContext, "tournament": tournament, "end": end, "reset": reset}
	logger.WithFields(u.Inject(ctx, b3.B3MultipleHeader)).WithFields(fields).Debug("")
	ctx, span := otel.Tracer(u.ModuleConfig.InstrumentationName).Start(
		ctx,
		"RegisterTournamentReset",
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
