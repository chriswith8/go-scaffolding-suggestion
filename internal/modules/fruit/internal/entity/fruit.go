package entity

import (
	"errors"
	"math/rand"

	"github.com/google/uuid"
)

type Fruit struct {
	ID     string      `json:"id" gorm:"primarykey"`
	Name   string      `json:"name"`
	Color  string      `json:"color"`
	Status FruitStatus `json:"status"`
}

func NewFruit(name string, color string) *Fruit {
	return &Fruit{
		ID:     uuid.New().String(),
		Name:   name,
		Color:  color,
		Status: Unripe,
	}
}

func (f *Fruit) Ripe() {
	f.Status = Ripe
}

func (f *Fruit) Rot() error {
	if f.IsRotten() {
		return errors.New("fruit is already rotten")
	}

	if rand.Intn(2) == 0 {
		return errors.New("fruit cannot be rotten")
	}
	f.Status = Rotten
	return nil
}

func (f *Fruit) Resurrect() {
	f.Status = Unripe
}

func (f *Fruit) Destroy() {
	f.Status = Destroyed
}

func (f *Fruit) IsRotten() bool {
	return f.Status == Rotten
}
