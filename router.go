package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

const StaticDir = "/site/static/"

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}
	router.
		PathPrefix(StaticDir).
		Handler(http.StripPrefix(StaticDir, http.FileServer(http.Dir("."+StaticDir))))
	return
}
