package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	dirToServe := filepath.Join("docs", "api")

	fs := http.FileServer(http.Dir(dirToServe))
	http.Handle("/", http.StripPrefix("/", fs))

	port := "8080"
	log.Printf("Serving files from %s on http://localhost:%s/", dirToServe, port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
