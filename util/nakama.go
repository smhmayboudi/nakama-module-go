package util

import (
	"context"

	"github.com/heroiclabs/nakama-common/runtime"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type NakamaContext struct {
	ClientIp       string            `json:"client_ip,omitempty" protobuf:"bytes,1,opt,name=client_ip,json=client_ip,proto3"`
	ClientPort     string            `json:"client_port,omitempty" protobuf:"bytes,2,opt,name=client_port,json=client_port,proto3"`
	Env            map[string]string `json:"env,omitempty" protobuf:"bytes,3,opt,name=env,json=env,proto3" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExecutionMode  string            `json:"execution_mode,omitempty" protobuf:"bytes,4,opt,name=execution_mode,json=execution_mode,proto3"`
	Headers        map[string]string `json:"headers,omitempty" protobuf:"bytes,5,opt,name=headers,json=headers,proto3" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Lang           string            `json:"lang,omitempty" protobuf:"bytes,6,opt,name=lang,json=lang,proto3"`
	MatchId        string            `json:"match_id,omitempty" protobuf:"bytes,7,opt,name=match_id,json=match_id,proto3"`
	MatchLabel     string            `json:"match_label,omitempty" protobuf:"bytes,8,opt,name=match_label,json=match_label,proto3"`
	MatchNode      string            `json:"match_node,omitempty" protobuf:"bytes,9,opt,name=match_node,json=match_node,proto3"`
	MatchTickRate  int               `json:"match_tick_rate,omitempty" protobuf:"bytes,10,opt,name=match_tick_rate,json=match_tick_rate,proto3"`
	Node           string            `json:"node,omitempty" protobuf:"bytes,11,opt,name=node,json=node,proto3"`
	QueryParams    map[string]string `json:"query_params,omitempty" protobuf:"bytes,12,opt,name=query_params,json=query_params,proto3" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SessionId      string            `json:"session_id,omitempty" protobuf:"bytes,13,opt,name=session_id,json=session_id,proto3"`
	UserId         string            `json:"user_id,omitempty" protobuf:"bytes,14,opt,name=user_id,json=user_id,proto3"`
	UserSessionExp int               `json:"user_session_exp,omitempty" protobuf:"bytes,15,opt,name=user_session_exp,json=user_session_exp,proto3"`
	Username       string            `json:"username,omitempty" protobuf:"bytes,16,opt,name=username,json=username,proto3"`
	Vars           map[string]string `json:"vars,omitempty" protobuf:"bytes,17,opt,name=vars,json=vars,proto3" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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
