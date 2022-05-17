package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:        ":8888",
		Handler:     routes(),
		IdleTimeout: time.Minute,
	}

	srv.ListenAndServe()
}

func routes() http.Handler {
	mux := http.NewServeMux()

	// index
	mux.HandleFunc("/", indexRoute)

	// static files
	httpFS := http.FileServer(http.Dir("ui/dist"))
	mux.Handle("/static/", httpFS)

	// api
	mux.HandleFunc("/api/v1/greeting", greetingAPI)
	return mux
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
		return
	}

	if r.URL.Path == "/favicon.ico" {
		http.ServeFile(w, r, "ui/dist/favicon.ico")
		return
	}

	http.ServeFile(w, r, "ui/dist/index.html")
}

func greetingAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, there!"))
}
