// https://github.com/heroiclabs/nakama/blob/master/sample_go_module/sample.go
package register

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type Match struct{}

type MatchState struct {
	debug bool
}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	ctx = trace.ContextWithRemoteSpanContext(ctx, u.NewSpanContext(ctx))
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"MatchInit",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterMatch.MatchInit", "params": params}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return params, 0, ""
	}
	var debug bool
	if d, ok := params["debug"]; ok {
		if dv, ok := d.(bool); ok {
			debug = dv
		}
	}
	state := &MatchState{
		debug: debug,
	}
	if state.debug {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).Info(fmt.Sprintf("match init, starting with debug: %v", state.debug))
	}
	tickRate := 1
	label := "skill=100-150"
	return state, tickRate, label
}
func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	ctx = trace.ContextWithRemoteSpanContext(ctx, u.NewSpanContext(ctx))
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"MatchJoinAttempt",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterMatch.MatchJoinAttempt", "tick": tick, "state": state, "presence": presence, "metadata": metadata}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return state, false, ""
	}
	if state.(*MatchState).debug {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).Info(fmt.Sprintf("match join attempt username %v user_id %v session_id %v node %v with metadata %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId(), metadata))
	}
	return state, true, ""
}
func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	ctx = trace.ContextWithRemoteSpanContext(ctx, u.NewSpanContext(ctx))
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"MatchJoin",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterMatch.MatchJoin", "tick": tick, "state": state, "presences": presences}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return state
	}
	if state.(*MatchState).debug {
		for _, presence := range presences {
			textMapCarrier := u.NewTextMapCarrier(ctx)
			logger.WithFields(textMapCarrier.MultipleField()).Info(fmt.Sprintf("match join username %v user_id %v session_id %v node %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId()))
		}
	}
	return state
}
func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	ctx = trace.ContextWithRemoteSpanContext(ctx, u.NewSpanContext(ctx))
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"MatchLeave",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterMatch.MatchLeave", "tick": tick, "state": state, "presences": presences}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return state
	}
	if state.(*MatchState).debug {
		for _, presence := range presences {
			textMapCarrier := u.NewTextMapCarrier(ctx)
			logger.WithFields(textMapCarrier.MultipleField()).Info(fmt.Sprintf("match leave username %v user_id %v session_id %v node %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId()))
		}
	}
	return state
}
func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	ctx = trace.ContextWithRemoteSpanContext(ctx, u.NewSpanContext(ctx))
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"MatchLoop",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterMatch.MatchLeave", "tick": tick, "state": state, "messages": messages}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return state
	}
	if state.(*MatchState).debug {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).Info(fmt.Sprintf("match loop match_id %v tick %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), tick))
		logger.WithFields(textMapCarrier.MultipleField()).Info(fmt.Sprintf("match loop match_id %v message count %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), len(messages)))
	}
	if tick >= 10 {
		return nil
	}
	return state
}
func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	ctx = trace.ContextWithRemoteSpanContext(ctx, u.NewSpanContext(ctx))
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"MatchTerminate",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterMatch.MatchTerminate", "tick": tick, "state": state, "graceSeconds": graceSeconds}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return state
	}
	if state.(*MatchState).debug {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).Info(fmt.Sprintf("match terminate match_id %v tick %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), tick))
		logger.WithFields(textMapCarrier.MultipleField()).Info(fmt.Sprintf("match terminate match_id %v grace seconds %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), graceSeconds))
	}
	return state
}
func (m *Match) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	ctx = trace.ContextWithRemoteSpanContext(ctx, u.NewSpanContext(ctx))
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"MatchSignal",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterMatch.MatchSignal", "tick": tick, "state": state, "data": data}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return state, data
	}
	if state.(*MatchState).debug {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.MultipleField()).Info(fmt.Sprintf("match signal match_id %v tick %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), tick))
		logger.WithFields(textMapCarrier.MultipleField()).Info(fmt.Sprintf("match signal match_id %v data %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), data))
	}
	return state, data
}

var _ runtime.Match = (*Match)(nil)

func RegisterMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
	ctx = trace.ContextWithRemoteSpanContext(ctx, u.NewSpanContext(ctx))
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterMatch",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	return &Match{}, nil
}
