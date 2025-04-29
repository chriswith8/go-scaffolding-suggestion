package web

import (
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/feature"
	pgValidator "github.com/go-playground/validator"
)

type (
	FruitHandler struct {
		fruit     *feature.Fruit
		validator *pgValidator.Validate
	}
)

func NewFruitHandler(fruit *feature.Fruit, validator *pgValidator.Validate) *FruitHandler {
	return &FruitHandler{fruit: fruit, validator: validator}
}
