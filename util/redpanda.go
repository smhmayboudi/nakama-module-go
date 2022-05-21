package util

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/heroiclabs/nakama-common/runtime"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type NakamaContext struct {
	ClientIp       string            `json:"client_ip,omitempty"`
	ClientSort     string            `json:"client_port,omitempty"`
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

type DataKey struct {
	Node string `json:"node,omitempty"`
}

type DataValue struct {
	B3            string                 `json:"b3,omitempty"`
	NakamaContext NakamaContext          `json:"context,omitempty"`
	Payload       map[string]interface{} `json:"payload,omitempty"`
}

type Record struct {
	Key       DataKey   `json:"key,omitempty"`
	Partition int       `json:"partition,omitempty"`
	Value     DataValue `json:"value,omitempty"`
}

type Records struct {
	Records []Record `json:"records,omitempty"`
}

func Redpanda(ctx context.Context, logger runtime.Logger, payloadValue map[string]interface{}) error {
	ctx, span := otel.Tracer(InstrumentationName).Start(
		ctx,
		"Redpanda",
		trace.WithSpanKind(trace.SpanKindProducer))
	defer span.End()

	clientIp, ok := ctx.Value(runtime.RUNTIME_CTX_CLIENT_IP).(string)
	if !ok {
		clientIp = ""
	}
	clientSort, ok := ctx.Value(runtime.RUNTIME_CTX_CLIENT_PORT).(string)
	if !ok {
		clientSort = ""
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
	nakamaContextValue := &NakamaContext{
		ClientIp:       clientIp,
		ClientSort:     clientSort,
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
	dataKey := &DataKey{
		Node: node,
	}
	b3Value := Inject(ctx, b3.B3SingleHeader)["b3"].(string)
	dataValue := &DataValue{
		B3:            b3Value,
		NakamaContext: *nakamaContextValue,
		Payload:       payloadValue,
	}
	record := &Record{
		Key:       *dataKey,
		Partition: 0,
		Value:     *dataValue,
	}
	items := make([]Record, 1)
	items[0] = *record
	records := &Records{
		Records: items,
	}
	body, err := json.Marshal(records)
	if err != nil {
		logger.WithFields(Inject(ctx, b3.B3MultipleHeader)).WithField("error", err).Error("Failed to marshaling to JSON")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to marshaling to JSON")
		return err
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", RedpandaURL, bytes.NewReader(body))
	if err != nil {
		logger.WithFields(Inject(ctx, b3.B3MultipleHeader)).WithField("error", err).Error("Failed to create request with context")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to create request with context")
		return err
	}
	req.Header.Add("Content-Type", "application/vnd.kafka.json.v2+json")
	for k, v := range Inject(ctx, b3.B3SingleHeader) {
		req.Header.Add(k, v.(string))
	}
	// client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		logger.WithFields(Inject(ctx, b3.B3MultipleHeader)).WithField("error", err).Error("Failed to create http client")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to create http client")
		return err
	}
	defer res.Body.Close()
	return nil
}
