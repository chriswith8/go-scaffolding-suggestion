package service

import "github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"

func (f fruitService) Rot(fruitInstance entity.Fruit) (*entity.Fruit, error) {
	if err := fruitInstance.Rot(); err != nil {
		return nil, err
	}

	if err := f.Repo.Save(fruitInstance); err != nil {
		return nil, err
	}
	return &fruitInstance, nil
}
