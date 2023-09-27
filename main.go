package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `jason:id`
	Title  string `jason:title`
	Author string `jason:author`
	Year   string `jason:year`
}

var books []Book

func main() {
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Title: "Book 1", Author: "Author 1", Year: "1991"},
		Book{ID: "2", Title: "Book 2", Author: "Author 2", Year: "1992"},
		Book{ID: "3", Title: "Book 3", Author: "Author 3", Year: "1993"},
		Book{ID: "4", Title: "Book 4", Author: "Author 4", Year: "1994"})

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
		}
	}
	json.NewEncoder(w).Encode([]string{})
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello World!")
}
