package fruit

import (
	"net/http"

	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/utils/validator"
	"github.com/chriswith8/go-scaffolding-suggestion/internal/web/controller"
	"github.com/gorilla/mux"
)

type (
	controllerImpl struct {
		handler *fruit.Handler
	}
)

func NewController(mdl *modules.Modules) controller.Controller {
	return &controllerImpl{handler: fruit.NewHandler(mdl.Fruit, validator.Validator())}
}

func (c *controllerImpl) BuildEndpoints(router *mux.Router) {
	router.HandleFunc("/fruits", c.handler.CreateFruit).Methods(http.MethodPost)
	router.HandleFunc("/fruits/{fruit_id}", c.handler.GetFruitByID).Methods(http.MethodGet)
	router.HandleFunc("/fruits/{fruit_id}/rot", c.handler.Rot).Methods(http.MethodPatch)
}
