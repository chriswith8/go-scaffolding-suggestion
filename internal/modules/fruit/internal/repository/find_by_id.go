package repository

import "github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"

func (r *repository) FindByID(id string) (*entity.Fruit, error) {
	var q = &entity.Fruit{ID: id}

	result := r.sqlAdapter.First(q)
	if result.Error != nil {
		return nil, result.Error
	}

	return q, nil
}
