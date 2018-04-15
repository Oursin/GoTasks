package router

import "net/http"
import "GoTasks/src/handlers"

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

var routesIndex = routes{
	route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	route{
		"Task",
		"GET",
		"/task/{id}",
		handlers.TasksHandler,
	},
}
