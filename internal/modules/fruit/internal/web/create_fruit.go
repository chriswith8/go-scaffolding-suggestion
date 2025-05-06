package web

import (
	"encoding/json"
	"net/http"

	"github.com/chriswith8/go-scaffolding-suggestion/internal/modules/fruit/internal/data"
)

func (h FruitHandler) CreateFruit(w http.ResponseWriter, r *http.Request) {
	var payload data.NewFruitInput

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.validator.StructCtx(r.Context(), payload); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	fruit, err := h.uc.CreateFruit.Exec(payload)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return

	}

	if err := json.NewEncoder(w).Encode(fruit); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
