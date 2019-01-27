// rest api server implementation in golang

package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct model

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	// loop through the books, find the correct ID
	for _, item := range books {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // mock id
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {

}
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// init the mux router
	r := mux.NewRouter()

	// mock data
	books = append(books, Book{ID: "1", Isbn: "1156454", Title: "Book 1", Author: &Author{Firstname: "John1", Lastname: "Doe1"}})
	books = append(books, Book{ID: "2", Isbn: "2156454", Title: "Book 2", Author: &Author{Firstname: "John2", Lastname: "Doe2"}})
	books = append(books, Book{ID: "3", Isbn: "3156454", Title: "Book 3", Author: &Author{Firstname: "John3", Lastname: "Doe3"}})

	// route handlers and endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
