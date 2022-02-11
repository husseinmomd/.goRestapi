package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// books stuct (MODEL)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

// get all book
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// create a new book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}
func main() {
	// init Router
	r := mux.NewRouter()

	// route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}
