package service

import (
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/data"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"
)

func (f fruitService) Create(payload data.NewFruitInput) (*entity.Fruit, error) {
	fruit := entity.NewFruit(payload.Name, payload.Color)
	if err := f.Repo.Create(*fruit); err != nil {
		return nil, err
	}
	return fruit, nil
}
