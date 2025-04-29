package web

import (
	"encoding/json"
	"net/http"

	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/data"
)

func (h FruitHandler) CreateFruit(w http.ResponseWriter, r *http.Request) {
	var payload data.NewFruitInput

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(payload); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	fruit, err := h.fruit.CreateFruit.Exec(payload)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return

	}

	if err := json.NewEncoder(w).Encode(fruit); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
