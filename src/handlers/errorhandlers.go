package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Handler404 defines the server's behaviour when the resource requested is not available.
func Handler404(w http.ResponseWriter, _ *http.Request) {
	base, err := os.Getwd()
	if err != nil {
		log.Fatal("There was an error getting current working directory")
	}
	path := base + HTMLPath + "404.html"
	file, _ := os.Open(path)
	if err != nil {
		log.Fatal("There was an error getting 404 html page")
	}
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
