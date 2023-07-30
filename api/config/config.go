package config

import (
	"os"
	"strings"
)

var (
	Name                      string
	NameForPrometheus         string
	Version                   string
	Env                       string
	BaseUrl                   string
	Port                      string
	OTelExporterHost          string
	OTelExporterPort          string
	OTelExporterUrlPath       string
	OTelExporterAuthorization string
)

func init() {
	set()
}

func set() {
	Name = os.Getenv("NAME")
	NameForPrometheus = strings.ReplaceAll(Name, "-", "_")
	Version = os.Getenv("VERSION")
	Env = os.Getenv("ENV")
	BaseUrl = os.Getenv("BASE_URL")
	Port = os.Getenv("PORT")
	OTelExporterHost = os.Getenv("OTEL_EXPORTER_HOST")
	OTelExporterPort = os.Getenv("OTEL_EXPORTER_PORT")
	OTelExporterUrlPath = os.Getenv("OTEL_EXPORTER_URL_PATH")
	OTelExporterAuthorization = os.Getenv("OTEL_EXPORTER_AUTHORIZATION")
}
