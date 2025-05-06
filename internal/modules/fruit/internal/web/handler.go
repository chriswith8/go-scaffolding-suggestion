package web

import (
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/usecase"
	pgValidator "github.com/go-playground/validator"
)

type (
	FruitHandler struct {
		uc        *usecase.UseCases
		validator *pgValidator.Validate
	}
)

func NewFruitHandler(useCases *usecase.UseCases, validator *pgValidator.Validate) *FruitHandler {
	return &FruitHandler{uc: useCases, validator: validator}
}
