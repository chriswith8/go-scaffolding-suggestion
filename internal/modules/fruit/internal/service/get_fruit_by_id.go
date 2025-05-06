package service

import "github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"

func (f fruitService) GetFruitByID(id string) (*entity.Fruit, error) {
	fruitInstance, err := f.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return fruitInstance, nil
}
