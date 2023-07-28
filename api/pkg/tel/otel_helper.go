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
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type OTelHttpConfig struct {
	ExporterConfig   OTelHttpExporterConfig
	AttributesConfig OTelHttpAttributesConfig
}

type OTelHttpExporterConfig struct {
	Host          string
	Port          string
	UrlPath       string
	Authorization string
	IsEnabledSSL  bool
}

type OTelHttpAttributesConfig struct {
	Name        string
	Version     string
	Environment string
}

func InitOTelHttpTracer(option OTelHttpConfig) (*sdktrace.TracerProvider, error) {
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	otlptracehttp.NewClient()

	otlpHttpExporter, err := newOtlpHttpExporter(option.ExporterConfig)
	if err != nil {
		return nil, err
	}

	traceProvider := newTraceProvider(option.AttributesConfig, otlpHttpExporter)

	return traceProvider, nil
}

func newOtlpHttpExporter(cfg OTelHttpExporterConfig) (*otlptrace.Exporter, error) {
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

func newTraceProvider(cfg OTelHttpAttributesConfig, exporter *otlptrace.Exporter) *sdktrace.TracerProvider {
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(cfg.Name),
		semconv.ServiceVersionKey.String(cfg.Version),
		attribute.String("environment", cfg.Environment),
	)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	return tp
}
