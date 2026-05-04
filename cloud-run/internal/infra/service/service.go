package service

import (
	"cepclima/internal/entity"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var (
	cepURL    = "https://viacep.com.br/ws/"
	cepURLEnd = "/json/"
	climaURL  = "http://api.weatherapi.com/v1/current.json"
	tracer    = otel.Tracer("cepclima/service")
)

type CepClimaService struct{}

func NewCepClimaService() *CepClimaService {
	return &CepClimaService{}
}

func (s *CepClimaService) GetCep(ctx context.Context, cepString string) (*entity.Cep, error) {
	ctx, span := tracer.Start(ctx, "viacep.GetCep")
	defer span.End()
	span.SetAttributes(attribute.String("cep", cepString))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cepURL+cepString+cepURLEnd, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cep entity.Cep
	if err = json.NewDecoder(resp.Body).Decode(&cep); err != nil || cep.Erro == "true" {
		return nil, errors.New("can not find zipcode")
	}

	return &cep, nil
}

func (s *CepClimaService) GetClimaByCep(ctx context.Context, cidade string) (*entity.ClimaOutput, error) {
	ctx, span := tracer.Start(ctx, "weatherapi.GetClima")
	defer span.End()
	span.SetAttributes(attribute.String("city", cidade))

	weatherAPI := os.Getenv("WEATHER_API")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, climaURL, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", weatherAPI)
	q.Add("q", cidade)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var clima entity.Clima
	json.NewDecoder(resp.Body).Decode(&clima)

	return &entity.ClimaOutput{
		Celsius:    clima.Current.Temp_c,
		Fahrenheit: clima.Current.Temp_f,
		Kelvin:     clima.Current.Temp_c + 273.15,
	}, nil
}
