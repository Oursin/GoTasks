package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const StaticDir = "/site/static/"

func NewRouter() (router *mux.Router) {
	basepath, err := os.Getwd()
	if err != nil {
		log.Fatal("There was an error loading current Working directory")
	}
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
		Handler(http.StripPrefix(StaticDir, http.FileServer(http.Dir(basepath+StaticDir))))
	return
}
