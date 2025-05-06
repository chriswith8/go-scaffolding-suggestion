package service

import (
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/data"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/repository"
)

type (
	FruitService interface {
		Create(data.NewFruitInput) (*entity.Fruit, error)
		Rot(entity.Fruit) (*entity.Fruit, error)
		Destroy(entity.Fruit) (*entity.Fruit, error)
		GetFruitByID(id string) (*entity.Fruit, error)
	}
	fruitService struct {
		Dependencies
	}
	Dependencies struct {
		Repo repository.Repository
	}
)

func NewFruitService(deps Dependencies) FruitService {
	return &fruitService{deps}
}
