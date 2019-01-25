// rest api server implementation in golang

package main

import (
	"github.com/gorilla/mux"
)

// Book struct model

type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {

}

func main() {
	// init the mux router
	r := mux.NewRouter()

	// route handlers and endpoints

	r.HandleFunc("/api/books", getBooks).Medthods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Medthods("GET")
	r.HandleFunc("/api/books", createBook).Medthods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Medthods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Medthods("DELETE")

	// start the server
	log.Fatal(http.ListenAndServer(":8000", r)
}
