package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"

	"github.com/hyorimitsu/hello-openobserve/api/config"
	"github.com/hyorimitsu/hello-openobserve/api/pkg/tel"
)

var tracer = otel.Tracer("github.com/hyorimitsu/hello-openobserve/api")

func main() {
	option := tel.OTelHttpConfig{
		ExporterConfig: tel.OTelHttpExporterConfig{
			Host:          config.OTelExporterHost,
			Port:          config.OTelExporterPort,
			UrlPath:       config.OTelExporterUrlPath,
			Authorization: config.OTelExporterAuthorization,
			IsEnabledSSL:  config.Env != "local",
		},
		AttributesConfig: tel.OTelHttpAttributesConfig{
			Name:        config.Name,
			Version:     config.Version,
			Environment: config.Env,
		},
	}

	tracerProvider, err := tel.InitOTelHttpTracer(option)
	if err != nil {
		fmt.Println("[Error] Unable to initialize OTel HTTP exporter")
		return
	}
	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			fmt.Println("[Error] Unable to shutting down tracer provider: ", err)
		}
	}()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(otelecho.Middleware(config.Name))

	e.GET(config.BaseUrl+"/hello", hello)

	e.Logger.Fatal(e.Start(":" + config.Port))
}

func hello(ctx echo.Context) error {
	c := ctx.Request().Context()

	_, span := tracer.Start(c, "hello")
	defer span.End()

	fmt.Println("[Info] Called API")
	return ctx.String(http.StatusOK, "Hello, World!")
}
