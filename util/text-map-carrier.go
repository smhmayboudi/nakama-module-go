package util

import (
	"context"

	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel/propagation"
)

const (
	b3ContextHeader      = "b3"
	b3ParentSpanIDHeader = "x-b3-parentspanid"
	b3SampledHeader      = "x-b3-sampled"
	b3SpanIDHeader       = "x-b3-spanid"
	b3TraceIDHeader      = "x-b3-traceid"
)

type TextMapCarrier struct {
	B3           string `json:"b3,omitempty"`
	ParentSpanId string `json:"parent_span_id,omitempty"`
	Sampled      string `json:"sampled,omitempty"`
	SpanId       string `json:"span_id,omitempty"`
	TraceId      string `json:"trace_id,omitempty"`
}

func (m *TextMapCarrier) Get(key string) string {
	switch key {
	case b3ContextHeader:
		return m.B3
	case b3ParentSpanIDHeader:
		return m.ParentSpanId
	case b3SampledHeader:
		return m.Sampled
	case b3SpanIDHeader:
		return m.SpanId
	case b3TraceIDHeader:
		return m.TraceId
	}
	return ""
}

func (m *TextMapCarrier) Set(key string, value string) {
	switch key {
	case b3ContextHeader:
		m.B3 = value
	case b3ParentSpanIDHeader:
		m.ParentSpanId = value
	case b3SampledHeader:
		m.Sampled = value
	case b3SpanIDHeader:
		m.SpanId = value
	case b3TraceIDHeader:
		m.TraceId = value
	}
}

func (m *TextMapCarrier) Keys() []string {
	return []string{
		b3ContextHeader,
		b3ParentSpanIDHeader,
		b3SampledHeader,
		b3SpanIDHeader,
		b3TraceIDHeader,
	}
}

var _ propagation.TextMapCarrier = (*TextMapCarrier)(nil)

func (m *TextMapCarrier) MultipleField() map[string]interface{} {
	field := map[string]interface{}{}
	if m.ParentSpanId != "" {
		field["parent_span_id"] = m.ParentSpanId
	}
	if m.Sampled != "" {
		field["sampled"] = m.Sampled
	}
	if m.SpanId != "" {
		field["span_id"] = m.SpanId
	}
	if m.TraceId != "" {
		field["trace_id"] = m.TraceId
	}
	return field
}

func (m *TextMapCarrier) SingleField() map[string]interface{} {
	field := map[string]interface{}{}
	if m.B3 != "" {
		field["b3"] = m.B3
	}
	return field
}

func NewTextMapCarrier(ctx context.Context) *TextMapCarrier {
	textMapCarrier := TextMapCarrier{}
	b3 := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader | b3.B3SingleHeader))
	b3.Inject(ctx, &textMapCarrier)
	return &textMapCarrier
}
