package main

import (
	"cepclima/internal/infra/controller"
	"cepclima/internal/infra/service"
	"cepclima/internal/infra/tracer"
	"cepclima/internal/usecase"
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	if err := godotenv.Load("cmd/.env"); err != nil {
		log.Printf("aviso: arquivo .env não encontrado, usando variáveis de ambiente")
	}

	endpoint := os.Getenv("OTEL_EXPORTER_ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:4318"
	}
	shutdown := tracer.Init("service-b", endpoint)
	defer shutdown(context.Background())

	svc := service.NewCepClimaService()
	uc := usecase.NewCepUsecase(svc)
	ctrl := controller.NewCepController(*uc)

	r := gin.Default()
	r.Use(otelgin.Middleware("service-b"))
	r.POST("/cep", ctrl.GetClima)
	r.Run(":8080")
}
