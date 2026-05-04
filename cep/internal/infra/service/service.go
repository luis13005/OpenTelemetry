package service

import (
	"bytes"
	"cep/internal/entity"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

var tracer = otel.Tracer("cep/service")

type CepService struct{}

func NewCepService() *CepService {
	return &CepService{}
}

func (s *CepService) ForwardToServiceB(ctx context.Context, cep string) (*entity.ClimaOutput, error) {
	ctx, span := tracer.Start(ctx, "service-b.ForwardCep")
	defer span.End()

	serviceBURL := os.Getenv("SERVICE_B_URL")
	if serviceBURL == "" {
		serviceBURL = "http://localhost:8080"
	}

	payload, err := json.Marshal(entity.InputCep{Cep: cep})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, serviceBURL+"/cep", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Injeta o contexto de trace para rastreamento distribuído
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var msg string
		json.NewDecoder(resp.Body).Decode(&msg)
		return nil, fmt.Errorf("%s", msg)
	}

	var out entity.ClimaOutput
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
