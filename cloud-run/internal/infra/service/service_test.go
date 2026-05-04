package service

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	godotenv.Load("../../../cmd/.env")
}

func TestGetCep(t *testing.T) {
	svc := NewCepClimaService()
	cep, err := svc.GetCep(context.Background(), "01310100")

	require.NoError(t, err)
	assert.Equal(t, "São Paulo", cep.Localidade)
	assert.Equal(t, "SP", cep.Uf)
}

func TestGetClimaByCep(t *testing.T) {
	svc := NewCepClimaService()
	out, err := svc.GetClimaByCep(context.Background(), "São Paulo")

	require.NoError(t, err)
	assert.NotEmpty(t, out.Celsius)
	assert.NotEmpty(t, out.Fahrenheit)
	assert.NotEmpty(t, out.Kelvin)
}
