package register

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	u "github.com/smhmayboudi/nakama-modules-go/util"
)

func RegisterMatchmakerMatched(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, entries []runtime.MatchmakerEntry) (string, error) {
	ctx = u.Extract(ctx, b3.B3SingleHeader)
	nakamaContext := u.NewContext(ctx, logger)
	fields := map[string]interface{}{"name": "RegisterMatchmakerMatched", "ctx": nakamaContext, "entries": entries}
	logger.WithFields(u.Inject(ctx, b3.B3MultipleHeader)).WithFields(fields).Debug("")
	ctx, span := otel.Tracer(u.ModuleConfig.InstrumentationName).Start(
		ctx,
		"RegisterMatchmakerMatched",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.RedpandaSend(ctx, logger, fields); err != nil {
		logger.WithFields(u.Inject(ctx, b3.B3MultipleHeader)).WithFields(fields).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return "", u.InternalError
	}
	return "", nil
}
