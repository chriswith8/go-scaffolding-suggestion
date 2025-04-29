package web

import (
	"fmt"
	"os"

	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/web/controller"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/web/controller/fruit"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/web/controller/ping"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (app Application) Run() error {
	dbConn, err := app.openDBConn()
	if err != nil {
		return err
	}
	m := modules.NewModules(dbConn)
	router := app.startRouter()

	app.startControllers(m, router)

	return nil
}

func (app *Application) startControllers(m *modules.Modules, router *mux.Router) {
	controllers := controller.Collection{
		fruit.NewController(m),
		ping.NewController(),
	}

	for _, ctrl := range controllers {
		ctrl.BuildEndpoints(router)
	}
}

func (app *Application) openDBConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func (app *Application) startRouter() *mux.Router {
	return mux.NewRouter()
}
