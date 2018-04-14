package router

import (
	"log"
	"net/http"
	"os"

	"GoTasks/src/handlers"

	"github.com/gorilla/mux"
)

const staticDir = "/site/static/"

func createStaticRouter(router *mux.Router) {
	basepath, err := os.Getwd()
	if err != nil {
		log.Fatal("There was an error loading current Working directory")
	}
	handler := http.StripPrefix(staticDir, http.FileServer(http.Dir(basepath+staticDir)))
	handler = checkPath(handler, "Files")
	router.
		PathPrefix(staticDir).
		Handler(handler)
}

func create404Router(router *mux.Router) {
	var handlerFunc http.HandlerFunc = handlers.Handler404
	router.NotFoundHandler = logger(handlerFunc, "404 Error")
}

// NewRouter initializes a new gorilla mux router
func NewRouter() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)
	for _, route := range routesIndex {
		var handler http.Handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	createStaticRouter(router)
	create404Router(router)
	return
}
