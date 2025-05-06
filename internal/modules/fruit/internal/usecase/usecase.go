package usecase

import (
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/service"
)

type (
	UseCases struct {
		CreateFruit  CreateFruitUseCase
		RotFruit     RotFruitUseCase
		GetFruitByID GetFruitByIDUseCase
	}
)

func NewUseCases(svc service.FruitService) *UseCases {
	return &UseCases{
		CreateFruit:  NewCreateFruitUseCase(svc),
		RotFruit:     NewRotFruitUseCase(svc),
		GetFruitByID: NewGetFruitByIDUseCase(svc),
	}
}
