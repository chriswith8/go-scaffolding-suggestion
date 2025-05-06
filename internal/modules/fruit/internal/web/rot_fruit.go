package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (h FruitHandler) Rot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fruitID := mux.Vars(r)["fruit_id"]
	fruit, err := h.uc.RotFruit.Exec(fruitID)
	if err != nil {
		if err := json.NewEncoder(w).Encode(map[string]string{"error": err.Error()}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}

	if err := json.NewEncoder(w).Encode(fruit); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
