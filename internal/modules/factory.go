package modules

import (
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit"
	"gorm.io/gorm"
)

type Modules struct {
	Fruit *fruit.Module
}

func NewModules(dbConn *gorm.DB) *Modules {
	defer func() {
		r := recover()
		if r != nil {
			panic(r)
		}
	}()

	return &Modules{
		Fruit: fruit.NewModule(dbConn),
	}
}
