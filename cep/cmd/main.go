package main

import (
	"cep/internal/infra/controller"
	"cep/internal/infra/service"
	"cep/internal/infra/tracer"
	"cep/internal/usecase"
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
	tracer.Init("service-a", endpoint)

	svc := service.NewCepService()
	uc := usecase.NewCepUsecase(svc)
	ctrl := controller.NewCepController(uc)

	r := gin.Default()
	r.Use(otelgin.Middleware("service-a"))
	r.POST("/cep", ctrl.GetClima)
	r.Run(":8080")
}
