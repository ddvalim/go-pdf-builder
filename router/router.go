package router

import (
	"github.com/ddvalim/go-pdf-builder/core/ports"
	"github.com/ddvalim/go-pdf-builder/router/routes"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter()

	var r []ports.Route

	r = append(r, routes.PDF...)

	for _, route := range r {
		router.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return router
}
