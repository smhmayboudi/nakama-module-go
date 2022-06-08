import { B3InjectEncoding, B3Propagator } from "@opentelemetry/propagator-b3";
import { BatchSpanProcessor, ConsoleSpanExporter, SimpleSpanProcessor } from "@opentelemetry/sdk-trace-base";
import { FetchInstrumentation } from "@opentelemetry/instrumentation-fetch";
import {
  OTEL_ATTRIBUTE_COUNT_LIMIT,
  OTEL_ATTRIBUTE_VALUE_LENGTH_LIMIT,
  OTEL_BSP_EXPORT_TIMEOUT,
  OTEL_BSP_MAX_EXPORT_BATCH_SIZE,
  OTEL_BSP_MAX_QUEUE_SIZE,
  OTEL_BSP_SCHEDULE_DELAY,
  OTEL_EXPORTER_OTLP_CONCURRENCY_LIMIT,
  OTEL_EXPORTER_OTLP_HOSTNAME,
  // OTEL_EXPORTER_OTLP_TRACES_COMPRESSION,
  OTEL_EXPORTER_OTLP_TRACES_ENDPOINT,
  OTEL_EXPORTER_OTLP_TRACES_HEADERS,
  OTEL_EXPORTER_OTLP_TRACES_TIMEOUT,
  OTEL_FORCE_FLUSH_TIMEOUT,
  OTEL_RESOURCE_ATTRIBUTES,
  OTEL_SERVICE_NAME,
  OTEL_SPAN_ATTRIBUTE_COUNT_LIMIT,
  OTEL_SPAN_ATTRIBUTE_VALUE_LENGTH_LIMIT,
  OTEL_SPAN_EVENT_COUNT_LIMIT,
  OTEL_SPAN_LINK_COUNT_LIMIT,
} from "./constant";
import { OTLPTraceExporter } from "@opentelemetry/exporter-trace-otlp-http";
import { OpentelemetryLogger } from "./opentelemetry-logger";
import { SemanticAttributes, SemanticResourceAttributes } from "@opentelemetry/semantic-conventions";
import { WebTracerProvider } from "@opentelemetry/sdk-trace-web";
import { XMLHttpRequestInstrumentation } from "@opentelemetry/instrumentation-xml-http-request";
import { ZoneContextManager } from "@opentelemetry/context-zone";
import { registerInstrumentations } from "@opentelemetry/instrumentation";
import { type Exception, type Span, context, diag, DiagLogLevel, SpanKind, SpanStatusCode, trace } from "@opentelemetry/api";
import { type FetchError } from "@opentelemetry/instrumentation-fetch/build/src/types";
import { type ResourceAttributes, Resource } from "@opentelemetry/resources";

diag.setLogger(OpentelemetryLogger, DiagLogLevel.DEBUG);

const instrumentations =
  typeof self === "undefined"
    ? []
    : [
      new FetchInstrumentation({
        applyCustomAttributesOnSpan: (
          span: Span,
          _request: Request | RequestInit,
          result: FetchError | Response
        ): void => {
          const { attributes } = span as any;

          if (attributes.component === "fetch") {
            span.updateName(
              `${String(attributes["http.method"])} ${String(
                attributes["http.url"]
              )}`
            );
          }

          if (typeof result.status !== "undefined" && result.status > 299) {
            span.setStatus({
              code: SpanStatusCode.ERROR,
              message: (result as FetchError).message,
            });
          }
        },
        clearTimingResources: true,
        ignoreUrls:
          process.env.NODE_ENV === "production"
            ? []
            : [/^\/_next\/static.*/iu],
        propagateTraceHeaderCorsUrls: [
          /http:\/\/127.0.0.1:8090\.*/iu,
          /http:\/\/localhost:8090\.*/iu,
        ],
      }),
      new XMLHttpRequestInstrumentation({
        propagateTraceHeaderCorsUrls: [
          /http:\/\/127.0.0.1:8090\.*/iu,
          /http:\/\/localhost:8090\.*/iu,
        ]
      }),
    ],
  tracerProvider = new WebTracerProvider({
    forceFlushTimeoutMillis: OTEL_FORCE_FLUSH_TIMEOUT,
    generalLimits: {
      attributeCountLimit: OTEL_ATTRIBUTE_COUNT_LIMIT,
      attributeValueLengthLimit: OTEL_ATTRIBUTE_VALUE_LENGTH_LIMIT,
    },
    // idGenerator
    resource: new Resource(
      Object.fromEntries(
        `${SemanticResourceAttributes.SERVICE_NAME}=${OTEL_SERVICE_NAME}${OTEL_RESOURCE_ATTRIBUTES === "" ? "" : `,${OTEL_RESOURCE_ATTRIBUTES}`
          }`
          .split(",")
          .map((value) => value.split("="))
      ) as ResourceAttributes
    ),
    // sampler
    spanLimits: {
      attributeCountLimit: OTEL_SPAN_ATTRIBUTE_COUNT_LIMIT,
      attributeValueLengthLimit: OTEL_SPAN_ATTRIBUTE_VALUE_LENGTH_LIMIT,
      eventCountLimit: OTEL_SPAN_EVENT_COUNT_LIMIT,
      linkCountLimit: OTEL_SPAN_LINK_COUNT_LIMIT,
    },
  });

tracerProvider.addSpanProcessor(
  new BatchSpanProcessor(
    new OTLPTraceExporter({
      // compression: OTEL_EXPORTER_OTLP_TRACES_COMPRESSION,
      concurrencyLimit: OTEL_EXPORTER_OTLP_CONCURRENCY_LIMIT,
      headers: Object.fromEntries(
        OTEL_EXPORTER_OTLP_TRACES_HEADERS.split(",").map((value) =>
          value.split("=")
        )
      ),
      hostname: OTEL_EXPORTER_OTLP_HOSTNAME,
      httpAgentOptions: {},
      keepAlive: true,
      timeoutMillis: OTEL_EXPORTER_OTLP_TRACES_TIMEOUT,
      url: OTEL_EXPORTER_OTLP_TRACES_ENDPOINT,
    }),
    {
      exportTimeoutMillis: OTEL_BSP_EXPORT_TIMEOUT,
      maxExportBatchSize: OTEL_BSP_MAX_EXPORT_BATCH_SIZE,
      maxQueueSize: OTEL_BSP_MAX_QUEUE_SIZE,
      scheduledDelayMillis: OTEL_BSP_SCHEDULE_DELAY,
    }
  )
);
tracerProvider.addSpanProcessor(
  new SimpleSpanProcessor(new ConsoleSpanExporter())
);
tracerProvider.register({
  contextManager: new ZoneContextManager(),
  propagator: new B3Propagator({
    injectEncoding: B3InjectEncoding.SINGLE_HEADER,
  }),
});
registerInstrumentations({
  instrumentations,
  // meterProvider,
  tracerProvider,
});

function traceSpan<F extends (...args: any) => ReturnType<F>>(
  name: string,
  func: F,
  funcName: string
): ReturnType<F> {
  diag.debug("traceSpan called");

  const isProduction = process.env.NODE_ENV === "production",
    span = trace
      .getTracer("client", "2022.05.22")
      .startSpan(name, {
        attributes: {
          [SemanticAttributes.CODE_FUNCTION]: funcName,
        },
        kind: SpanKind.CLIENT,
        // links: {},
        // root: false,
        // startTime: new Date().getTime(),
      });

  if (isProduction) {
    return func();
  }

  return context.with(trace.setSpan(context.active(), span), () => {
    try {
      const result = func();

      return result;
    } catch (error) {
      span.recordException(error as Exception);
      span.setStatus({ code: SpanStatusCode.ERROR });
      throw error;
    } finally {
      span.end();
    }
  });
}

export { traceSpan };
