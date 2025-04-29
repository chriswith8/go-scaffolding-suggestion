package repository

import (
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"
)

func (r *repository) Create(fruit entity.Fruit) error {
	if res := r.sqlAdapter.Create(fruit); res.Error != nil {
		return res.Error
	}
	return nil
}
