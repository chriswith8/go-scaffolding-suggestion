package entity

import "gorm.io/gorm"

type Fruit struct {
	gorm.Model
	Name   string      `json:"name"`
	Color  string      `json:"color"`
	Status FruitStatus `json:"status"`
}

func NewFruit(name string, color string) *Fruit {
	return &Fruit{
		Name:   name,
		Color:  color,
		Status: Unripe,
	}
}

func (f *Fruit) Ripe() {
	f.Status = Ripe
}

func (f *Fruit) Rot() {
	f.Status = Rotten
}

func (f *Fruit) Resurrect() {
	f.Status = Unripe
}
