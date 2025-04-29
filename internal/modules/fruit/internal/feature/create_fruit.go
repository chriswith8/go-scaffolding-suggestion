package feature

import (
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/data"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/repository"
)

type (
	CreateFruitUseCase interface {
		Exec(payload data.NewFruitInput) (*entity.Fruit, error)
	}
	createFruitUseCase struct {
		repo repository.Repository
	}
)

func NewCreateFruitUseCase(repo repository.Repository) CreateFruitUseCase {
	return &createFruitUseCase{repo: repo}
}

func (c createFruitUseCase) Exec(payload data.NewFruitInput) (*entity.Fruit, error) {
	fruit := entity.NewFruit(payload.Name, payload.Color)
	if err := c.repo.Create(*fruit); err != nil {
		return nil, err
	}
	return fruit, nil
}
