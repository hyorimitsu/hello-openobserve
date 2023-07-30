package tel

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type OTelConfig struct {
	ExporterConfig   OTelExporterConfig
	AttributesConfig OTelAttributesConfig
}

type OTelExporterConfig struct {
	Host          string
	Port          string
	UrlPath       string
	Authorization string
	IsEnabledSSL  bool
}

type OTelAttributesConfig struct {
	Name        string
	Version     string
	Environment string
}

func InitOTelTracer(cfg OTelConfig) (*trace.TracerProvider, error) {
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	exporter, err := newExporter(cfg.ExporterConfig)
	if err != nil {
		return nil, err
	}

	traceProvider := newTracerProvider(cfg.AttributesConfig, exporter)
	otel.SetTracerProvider(traceProvider)

	return traceProvider, nil
}

func newExporter(cfg OTelExporterConfig) (*otlptrace.Exporter, error) {
	option := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)),
		otlptracehttp.WithURLPath(cfg.UrlPath),
		otlptracehttp.WithHeaders(map[string]string{
			"Authorization": cfg.Authorization,
		}),
	}
	if !cfg.IsEnabledSSL {
		option = append(option, otlptracehttp.WithInsecure())
	}

	exporter, err := otlptracehttp.New(context.Background(), option...)
	if err != nil {
		return nil, err
	}
	return exporter, nil
}

func newTracerProvider(cfg OTelAttributesConfig, exporter *otlptrace.Exporter) *trace.TracerProvider {
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(cfg.Name),
		semconv.ServiceVersionKey.String(cfg.Version),
		attribute.String("environment", cfg.Environment),
	)
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(res),
		trace.WithBatcher(exporter),
	)
	return tp
}
