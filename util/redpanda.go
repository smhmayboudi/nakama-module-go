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
	Node string `json:"node,omitempty" protobuf:"bytes,1,opt,name=node,json=node,proto3"`
}

type Value struct {
	B3      string                 `json:"b3,omitempty" protobuf:"bytes,1,opt,name=b3,json=b3,proto3"`
	Payload map[string]interface{} `json:"payload,omitempty" protobuf:"bytes,2,rep,name=vars,json=vars,proto3" protobuf_key:"bytes,1,opt,name=key,json=key,proto3" protobuf_val:"bytes,2,opt,name=value,json=value,proto3"`
}

type Record struct {
	Key       Key   `json:"key,omitempty" protobuf:"bytes,1,opt,name=key,json=key,proto3"`
	Partition int   `json:"partition,omitempty" protobuf:"varint,2,opt,name=partition,json=partition,proto3"`
	Value     Value `json:"value,omitempty" protobuf:"bytes,3,opt,name=value,json=value,proto3"`
}

type Records struct {
	Records []Record `json:"records,omitempty" protobuf:"bytes,1,rep,name=records,json=records,proto3"`
}

func RedpandaSend(ctx context.Context, logger runtime.Logger, payload map[string]interface{}) error {
	nakamaContext := NewContext(ctx, logger)
	fields := map[string]interface{}{"name": "RedpandaSend", "ctx": nakamaContext, "payload": payload}
	logger.WithFields(Inject(ctx, b3.B3MultipleHeader)).WithFields(fields).Debug("")
	ctx, span := otel.Tracer(ModuleConfig.InstrumentationName).Start(
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
		logger.WithFields(Inject(ctx, b3.B3MultipleHeader)).WithFields(fields).WithField("error", err).Error("Failed to marshaling to JSON")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to marshaling to JSON")
		return err
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", ModuleConfig.RedpandaURL, bytes.NewReader(body))
	if err != nil {
		logger.WithFields(Inject(ctx, b3.B3MultipleHeader)).WithFields(fields).WithField("error", err).Error("Failed to create request with context")
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
		logger.WithFields(Inject(ctx, b3.B3MultipleHeader)).WithFields(fields).WithField("error", err).Error("Failed to create http client")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to create http client")
		return err
	}
	defer res.Body.Close()
	return nil
}
