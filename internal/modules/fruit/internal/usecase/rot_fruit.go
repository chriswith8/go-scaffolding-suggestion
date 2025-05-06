package usecase

import (
	"fmt"
	"log"

	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/service"
)

type (
	RotFruitUseCase interface {
		Exec(fruitID string) (*entity.Fruit, error)
	}
	rotFruitUseCase struct {
		service service.FruitService
	}
)

func NewRotFruitUseCase(s service.FruitService) RotFruitUseCase {
	return &rotFruitUseCase{service: s}
}

func (r rotFruitUseCase) Exec(fruitID string) (*entity.Fruit, error) {
	f, err := r.service.GetFruitByID(fruitID)
	if err != nil {
		log.Println("error finding fruit:", err)
		return nil, err
	}

	updatedFruit, err := r.service.Rot(*f)
	if err != nil {
		if err.Error() == "fruit is already rotten" {
			return nil, err
		}

		log.Println(fmt.Errorf("error rotating fruit: %w. Fruit was destroyed", err))
		return r.destroy(*f)
	}

	return updatedFruit, nil
}

func (r rotFruitUseCase) destroy(f entity.Fruit) (*entity.Fruit, error) {
	destroyedFruit, err := r.service.Destroy(f)
	if err != nil {
		log.Println("error destroying fruit:", err)
		return nil, err
	}
	return destroyedFruit, nil
}
