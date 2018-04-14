package router

import (
	"log"
	"net/http"
	"time"
)

func isForbidden(path string) bool {
	switch {
	case path == "/site/static/", path == "/site/static/styles", path == "/site/static/js", path == "/site/static/images":
		return true
	default:
		return false
	}
}

func checkPath(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		if isForbidden(r.RequestURI) {
			http.Error(w, "<div style=\"text-align=center, width=100%\"<h1>Forbidden</h1></div>", http.StatusForbidden)
		}
		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			name,
			time.Since(start),
		)
	})
}
