package controller

import "github.com/gorilla/mux"

type (
	Controller interface {
		BuildEndpoints(router *mux.Router)
	}
	Collection []Controller
)
