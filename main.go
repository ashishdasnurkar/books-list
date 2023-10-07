package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type Book struct {
	ID     string `jason:id`
	Title  string `jason:title`
	Author string `jason:author`
	Year   string `jason:year`
}

var books []Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println(os.Getenv("PORT"))
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Title: "Book 1", Author: "Author 1", Year: "1991"},
		Book{ID: "2", Title: "Book 2", Author: "Author 2", Year: "1992"},
		Book{ID: "3", Title: "Book 3", Author: "Author 3", Year: "1993"},
		Book{ID: "4", Title: "Book 4", Author: "Author 4", Year: "1994"})

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", addBooks).Methods("POST")
	r.HandleFunc("/books", updateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", removeBook).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for i, book := range books {
		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)
}

func addBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
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
