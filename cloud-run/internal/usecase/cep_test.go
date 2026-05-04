package usecase_test

import (
	"cepclima/internal/infra/service"
	"cepclima/internal/usecase"
	"context"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	godotenv.Load("../../cmd/.env")
}

func TestGetCepClima(t *testing.T) {
	svc := service.NewCepClimaService()
	uc := usecase.NewCepUsecase(svc)

	out, err := uc.GetCepClima(context.Background(), "01310100")

	require.NoError(t, err)
	assert.Equal(t, "São Paulo", out.City)
	assert.NotEmpty(t, out.Celsius)
	assert.NotEmpty(t, out.Fahrenheit)
	assert.NotEmpty(t, out.Kelvin)
}
