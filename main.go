package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/books", getBooks).Methods("GET")

	//log.Fatal(http.ListenAndServe(":8080", router))
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("GENERAL KENOBI!")
}
