package usecase

import (
	"cep/internal/entity"
	"context"
)

type ServiceInterface interface {
	ForwardToServiceB(ctx context.Context, cep string) (*entity.ClimaOutput, error)
}

type CepUsecase struct {
	service ServiceInterface
}

func NewCepUsecase(svc ServiceInterface) *CepUsecase {
	return &CepUsecase{service: svc}
}

func (uc *CepUsecase) GetCepClima(ctx context.Context, cep string) (*entity.ClimaOutput, error) {
	return uc.service.ForwardToServiceB(ctx, cep)
}
