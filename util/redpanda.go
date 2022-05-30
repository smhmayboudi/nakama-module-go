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

type Key struct {
	Node string `json:"node,omitempty"`
}

type Value struct {
	B3      string                 `json:"b3,omitempty"`
	Payload map[string]interface{} `json:"payload,omitempty"`
}

type Record struct {
	Key       Key   `json:"key,omitempty"`
	Partition int   `json:"partition,omitempty"`
	Value     Value `json:"value,omitempty"`
}

type Records struct {
	Records []Record `json:"records,omitempty"`
}

func RedpandaSend(ctx context.Context, logger runtime.Logger, payload map[string]interface{}) error {
	nakamaContext := NewContext(ctx, logger)
	logger.WithFields(Inject(ctx, b3.B3MultipleHeader)).WithFields(map[string]interface{}{"name": "RedpandaSend", "ctx": nakamaContext, "payload": payload}).Debug("")
	ctx, span := otel.Tracer(AppConfig.InstrumentationName).Start(
		ctx,
		"RedpandaSend",
		trace.WithSpanKind(trace.SpanKindProducer))
	defer span.End()

	key := &Key{
		Node: nakamaContext.Node,
	}
	b3V := Inject(ctx, b3.B3SingleHeader)["b3"].(string)
	value := &Value{
		B3:      b3V,
		Payload: payload,
	}
	record := &Record{
		Key:   *key,
		Value: *value,
	}
	items := make([]Record, 1)
	items[0] = *record
	records := &Records{
		Records: items,
	}
	body, err := json.Marshal(records)
	if err != nil {
		logger.WithFields(Inject(ctx, b3.B3MultipleHeader)).WithFields(map[string]interface{}{"name": "RedpandaSend", "ctx": nakamaContext}).WithField("error", err).Error("Failed to marshaling to JSON")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to marshaling to JSON")
		return err
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", AppConfig.RedpandaURL, bytes.NewReader(body))
	if err != nil {
		logger.WithFields(Inject(ctx, b3.B3MultipleHeader)).WithFields(map[string]interface{}{"name": "RedpandaSend", "ctx": nakamaContext}).WithField("error", err).Error("Failed to create request with context")
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
		logger.WithFields(Inject(ctx, b3.B3MultipleHeader)).WithFields(map[string]interface{}{"name": "RedpandaSend", "ctx": nakamaContext}).WithField("error", err).Error("Failed to create http client")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to create http client")
		return err
	}
	defer res.Body.Close()
	return nil
}
