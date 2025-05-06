package fruit

import (
	"github.com/chriswith8/go-scaffolding-suggestion/internal/adapters/sql"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/entity"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/repository"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/service"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/usecase"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/web"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type (
	Module  = usecase.UseCases
	Handler = web.FruitHandler
	Fruit   = entity.Fruit
)

func NewModule(dbConn *gorm.DB) *Module {
	sqlAdapter := sql.NewSQLAdapter[entity.Fruit](dbConn)
	s := service.NewFruitService(service.Dependencies{
		Repo: repository.NewRepository(sqlAdapter),
	})

	return usecase.NewUseCases(s)
}

func NewHandler(module *Module, vld *validator.Validate) *Handler {
	return web.NewFruitHandler(module, vld)
}
