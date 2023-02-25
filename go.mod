module github.com/smhmayboudi/nakama-modules-go

go 1.18

require (
	github.com/heroiclabs/nakama-common v1.23.0
	github.com/kelseyhightower/envconfig v1.4.0
	go.opentelemetry.io/contrib/propagators/b3 v1.7.0
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/exporters/jaeger v1.7.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.7.0
	go.opentelemetry.io/otel/sdk v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	golang.org/x/sys v0.1.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
