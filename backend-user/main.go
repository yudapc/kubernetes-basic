package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Backend User!</h1><br>Visit the /api/users"))
}

func apiUsersHandler(w http.ResponseWriter, r *http.Request) {
	articles := []map[string]string{
		{
			"id":        "1",
			"full_name": "Jhon Dhoe",
		},
		{
			"id":        "2",
			"full_name": "Alex Martin",
		},
		{
			"id":        "3",
			"full_name": "Thomas Franky",
		},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(articles)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/users", apiUsersHandler)
	http.ListenAndServe(":"+port, mux)
}
