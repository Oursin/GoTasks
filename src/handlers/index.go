package handlers

import (
	"net/http"
	"os"
	"log"
	"io"
	"fmt"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	base, _ := os.Getwd()
	path := base + HtmlPath + "index.html"
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		log.Print("Error loading index.html")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	io.Copy(w, file)
}
