package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TODO: See if I can remove the method from the router and add it to the api function
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
