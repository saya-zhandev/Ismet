package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/projects", corsMiddleware(handlers.GetProject))
	http.HandleFunc("/api/projects/create", corsMiddleware(handlers.CreateProject))
	http.HandleFunc("/api/projects/delete", corsMiddleware(handlers.DeleteProject))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next(w, r)
	}
}
