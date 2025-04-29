package data

import "github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"

type NewFruitResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	Status entity.FruitStatus
}
