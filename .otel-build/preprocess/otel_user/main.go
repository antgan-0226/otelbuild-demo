package main

import _ "otelbuild-demo/otel_rules"

import _ "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
import _ "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
import _ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
import _ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
import _ "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
import _ "go.opentelemetry.io/otel"

import _ "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
import _ "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
import _ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
import _ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
import _ "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
import _ "go.opentelemetry.io/otel/sdk"
import _ "go.opentelemetry.io/otel"

import _ "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
import _ "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
import _ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
import _ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
import _ "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
import _ "go.opentelemetry.io/otel"

import (
	"context"
	"fmt"
	"net/http"
	"otelbuild-demo/service"
)

func main() {
	//测试net/http::RoundTrip的hook效果
	req, err := http.NewRequestWithContext(context.Background(), "GET", "http://www.baidu.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("otelbuild", "true")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	//测试service::CreateUser的hook效果
	u, _ := service.CreateUser(context.Background(), 1, "test")

	//model.UserModel注入了新字段age
	fmt.Println(fmt.Sprintf("UserModel: %+v", u))
}
