package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules"
	fruitModule "github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/utils/env"
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

	if err := app.migrateDB(dbConn); err != nil {
		return err
	}

	m := modules.NewModules(dbConn)
	router := app.startRouter()

	app.startControllers(m, router)

	log.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	return nil
}

func (app Application) startControllers(m *modules.Modules, router *mux.Router) {
	controllers := controller.Collection{
		fruit.NewController(m),
		ping.NewController(),
	}

	for _, ctrl := range controllers {
		ctrl.BuildEndpoints(router)
	}
}

func (app Application) openDBConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.GetVarWithDefault("DB_USER", "myuser"),
		env.GetVarWithDefault("DB_PASSWORD", "mypassword"),
		env.GetVarWithDefault("DB_HOST", "localhost"),
		env.GetVarWithDefault("DB_PORT", "3306"),
		env.GetVarWithDefault("DB_NAME", "mydatabase"),
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func (app Application) startRouter() *mux.Router {
	return mux.NewRouter()
}

func (app Application) migrateDB(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(
		fruitModule.Fruit{},
	)
}
