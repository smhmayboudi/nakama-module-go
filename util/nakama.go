package util

import (
	"context"

	"github.com/heroiclabs/nakama-common/runtime"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type NakamaContext struct {
	ClientIp       string            `json:"client_ip,omitempty"`
	ClientPort     string            `json:"client_port,omitempty"`
	Env            map[string]string `json:"env,omitempty"`
	ExecutionMode  string            `json:"execution_mode,omitempty"`
	Headers        map[string]string `json:"headers,omitempty"`
	Lang           string            `json:"lang,omitempty"`
	MatchId        string            `json:"match_id,omitempty"`
	MatchLabel     string            `json:"match_label,omitempty"`
	MatchNode      string            `json:"match_node,omitempty"`
	MatchTickRate  int               `json:"match_tick_rate,omitempty"`
	Node           string            `json:"node,omitempty"`
	QueryParams    map[string]string `json:"query_params,omitempty"`
	SessionId      string            `json:"session_id,omitempty"`
	UserId         string            `json:"user_id,omitempty"`
	UserSessionExp int               `json:"user_session_exp,omitempty"`
	Username       string            `json:"username,omitempty"`
	Vars           map[string]string `json:"vars,omitempty"`
}

func NewContext(ctx context.Context, logger runtime.Logger) *NakamaContext {
	ctx, span := otel.Tracer(ModuleConfig.InstrumentationName).Start(
		ctx,
		"NewContext",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	clientIp, ok := ctx.Value(runtime.RUNTIME_CTX_CLIENT_IP).(string)
	if !ok {
		clientIp = ""
	}
	clientPort, ok := ctx.Value(runtime.RUNTIME_CTX_CLIENT_PORT).(string)
	if !ok {
		clientPort = ""
	}
	env, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		env = map[string]string{}
	}
	executionMode, ok := ctx.Value(runtime.RUNTIME_CTX_MODE).(string)
	if !ok {
		executionMode = ""
	}
	headers, ok := ctx.Value(runtime.RUNTIME_CTX_HEADERS).(map[string]string)
	if !ok {
		headers = map[string]string{}
	}
	lang, ok := ctx.Value(runtime.RUNTIME_CTX_LANG).(string)
	if !ok {
		lang = ""
	}
	matchId, ok := ctx.Value(runtime.RUNTIME_CTX_MATCH_ID).(string)
	if !ok {
		matchId = ""
	}
	matchLabel, ok := ctx.Value(runtime.RUNTIME_CTX_MATCH_LABEL).(string)
	if !ok {
		matchLabel = ""
	}
	matchNode, ok := ctx.Value(runtime.RUNTIME_CTX_MATCH_NODE).(string)
	if !ok {
		matchNode = ""
	}
	matchTickRate, ok := ctx.Value(runtime.RUNTIME_CTX_MATCH_TICK_RATE).(int)
	if !ok {
		matchTickRate = 0
	}
	node, ok := ctx.Value(runtime.RUNTIME_CTX_NODE).(string)
	if !ok {
		node = ""
	}
	queryParams, ok := ctx.Value(runtime.RUNTIME_CTX_QUERY_PARAMS).(map[string]string)
	if !ok {
		queryParams = map[string]string{}
	}
	sessionId, ok := ctx.Value(runtime.RUNTIME_CTX_SESSION_ID).(string)
	if !ok {
		sessionId = ""
	}
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		userId = ""
	}
	userSessionExp, ok := ctx.Value(runtime.RUNTIME_CTX_USER_SESSION_EXP).(int)
	if !ok {
		userSessionExp = 0
	}
	username, ok := ctx.Value(runtime.RUNTIME_CTX_USERNAME).(string)
	if !ok {
		userId = ""
	}
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_VARS).(map[string]string)
	if !ok {
		vars = map[string]string{}
	}
	nakamaContext := &NakamaContext{
		ClientIp:       clientIp,
		ClientPort:     clientPort,
		Env:            env,
		ExecutionMode:  executionMode,
		Headers:        headers,
		Lang:           lang,
		MatchId:        matchId,
		MatchLabel:     matchLabel,
		MatchNode:      matchNode,
		MatchTickRate:  matchTickRate,
		Node:           node,
		QueryParams:    queryParams,
		SessionId:      sessionId,
		UserId:         userId,
		UserSessionExp: userSessionExp,
		Username:       username,
		Vars:           vars,
	}
	fields := map[string]interface{}{"name": "NewContext", "ctx": nakamaContext}
	logger.WithFields(Inject(ctx, b3.B3MultipleHeader)).WithFields(fields).Debug("")
	return nakamaContext
}
