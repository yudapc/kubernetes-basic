package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Category Page</h1>"))
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Article Page</h1>"))
}

func apiCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	articles := []map[string]string{
		{
			"id":    "1",
			"title": "Makanan Ringan",
		},
		{
			"id":    "2",
			"title": "Komputer & Gadget",
		},
		{
			"id":    "3",
			"title": "Finansial",
		},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(articles)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/category", categoryHandler)
	mux.HandleFunc("/article", articleHandler)
	mux.HandleFunc("/api/categories", apiCategoriesHandler)
	http.ListenAndServe(":"+port, mux)
}
