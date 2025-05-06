package usecase

import (
	"log"

	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/data"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/service"
)

type (
	CreateFruitUseCase interface {
		Exec(data.NewFruitInput) (*entity.Fruit, error)
	}
	createFruitUseCase struct {
		service service.FruitService
	}
)

func NewCreateFruitUseCase(s service.FruitService) CreateFruitUseCase {
	return &createFruitUseCase{service: s}
}

func (c createFruitUseCase) Exec(payload data.NewFruitInput) (*entity.Fruit, error) {
	f, err := c.service.Create(payload)

	if err != nil {
		log.Println("error creating fruit:", err)
		return nil, err
	}

	return f, nil
}
