package main

import "net/http"
import "GoTasks/handlers"

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route {
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
}
