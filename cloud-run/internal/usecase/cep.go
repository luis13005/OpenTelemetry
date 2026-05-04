package usecase

import (
	"cepclima/internal/entity"
	"context"
	"errors"
)

type ExternalServiceInterface interface {
	GetCep(ctx context.Context, cep string) (*entity.Cep, error)
	GetClimaByCep(ctx context.Context, cidade string) (*entity.ClimaOutput, error)
}

type UseCaseInterface interface {
	GetCepClima(ctx context.Context, cep string) (*entity.ClimaOutput, error)
}

type CepUsecase struct {
	service ExternalServiceInterface
}

func NewCepUsecase(service ExternalServiceInterface) *CepUsecase {
	return &CepUsecase{service: service}
}

func (uc *CepUsecase) GetCepClima(ctx context.Context, cepString string) (*entity.ClimaOutput, error) {
	cep, err := uc.service.GetCep(ctx, cepString)
	if err != nil {
		return nil, errors.New("can not find zipcode")
	}

	climaOut, err := uc.service.GetClimaByCep(ctx, cep.Localidade)
	if err != nil {
		return nil, err
	}

	climaOut.City = cep.Localidade
	return climaOut, nil
}
