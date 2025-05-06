package repository

import (
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"
)

func (r repository) Save(e entity.Fruit) error {
	if res := r.sqlAdapter.Save(e); res.Error != nil {
		return res.Error
	}
	return nil
}
