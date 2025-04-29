package fruit

import (
	"github.com/chriswith8/go-scaffolding-suggestion/internal/adapters/sql"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/feature"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/repository"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/web"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type (
	Fruit   = feature.Fruit
	Handler = web.FruitHandler
)

func NewFruit(dbConn *gorm.DB) *Fruit {
	sqlAdapter := sql.NewSQLAdapter[entity.Fruit](dbConn)

	return feature.NewFruitUseCase(feature.Dependencies{
		Repo: repository.NewRepository(sqlAdapter),
	})
}

func NewHandler(fruit *Fruit, vld *validator.Validate) *Handler {
	return web.NewFruitHandler(fruit, vld)
}
