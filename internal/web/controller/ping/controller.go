package ping

import (
	"encoding/json"
	"net/http"

	"github.com/chriswith8/go-scaffolding-suggestion/internal/web/controller"
	"github.com/gorilla/mux"
)

type controllerImpl struct{}

func NewController() controller.Controller {
	return &controllerImpl{}
}

func (c controllerImpl) BuildEndpoints(router *mux.Router) {
	router.PathPrefix("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode("pong"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
