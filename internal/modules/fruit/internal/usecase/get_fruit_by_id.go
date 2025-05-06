package usecase

import (
	"log"

	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/service"
)

type (
	GetFruitByIDUseCase interface {
		Exec(fruitID string) (*entity.Fruit, error)
	}

	getFruitByIDUseCase struct {
		service service.FruitService
	}
)

func NewGetFruitByIDUseCase(s service.FruitService) GetFruitByIDUseCase {
	return &getFruitByIDUseCase{service: s}
}

func (g getFruitByIDUseCase) Exec(fruitID string) (*entity.Fruit, error) {
	f, err := g.service.GetFruitByID(fruitID)
	if err != nil {
		log.Println("error finding fruit:", err)
		return nil, err
	}

	return f, nil
}
