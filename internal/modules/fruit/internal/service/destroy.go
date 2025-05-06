package service

import "github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"

func (f fruitService) Destroy(fruitInstance entity.Fruit) (*entity.Fruit, error) {
	fruitInstance.Destroy()
	if err := f.Repo.Save(fruitInstance); err != nil {
		return nil, err
	}
	return &fruitInstance, nil
}
