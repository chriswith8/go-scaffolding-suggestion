package data

type NewFruitInput struct {
	Name  string `json:"name" validate:"required"`
	Color string `json:"color" validate:"required"`
}
