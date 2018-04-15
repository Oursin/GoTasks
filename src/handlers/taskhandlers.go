package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
)

// TasksHandler handles simple task display.
func TasksHandler(w http.ResponseWriter, _ *http.Request) {
	base, _ := os.Getwd()
	path := base + TemplatesPath + "task.gohtml"
	file, err := os.Open(path)
	if err != nil {
		log.Print("Error loading index.html")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
