package repository

import (
	"github.com/chriswith8/go-scaffolding-suggestion/internal/adapters/sql"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"
)

type (
	Repository interface {
		Create(fruit entity.Fruit) error
		Save(e entity.Fruit) error
		FindByID(id string) (*entity.Fruit, error)
	}
	repository struct {
		sqlAdapter sql.Adapter[entity.Fruit]
	}
)

func NewRepository(sqlAdapter sql.Adapter[entity.Fruit]) Repository {
	return &repository{
		sqlAdapter: sqlAdapter,
	}
}
