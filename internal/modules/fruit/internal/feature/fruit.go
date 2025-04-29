package feature

import "github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/repository"

type (
	Fruit struct {
		CreateFruit CreateFruitUseCase
	}
	Dependencies struct {
		Repo repository.Repository
	}
)

func NewFruitUseCase(deps Dependencies) *Fruit {
	return &Fruit{
		CreateFruit: NewCreateFruitUseCase(deps.Repo),
	}
}
