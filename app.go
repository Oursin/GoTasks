package main

import (
	"GoTasks/src/router"
	"log"
	"net/http"
)

func main() {
	routermux := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8888", routermux))
}
