package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ashishdasnurkar/books-list/controllers"
	"github.com/ashishdasnurkar/books-list/driver"
	"github.com/ashishdasnurkar/books-list/models"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book
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
	db := driver.ConnectDB()
	controller := controllers.Controller{}
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	r.HandleFunc("/book/{id}", controller.GetBook(db)).Methods("GET")
	r.HandleFunc("/books", addBooks).Methods("POST")
	r.HandleFunc("/books", updateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", removeBook).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Exec("delete from books where id=$1", params["id"])
	logFatal(err)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)
	json.NewEncoder(w).Encode(rowsUpdated)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	log.Println(book)

	results, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)
	log.Println(err)
	log.Println(results)
	rowsUpdated, err := results.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)
}

func addBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var bookID int
	json.NewDecoder(r.Body).Decode(&book)
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id",
		book.Title, book.Author, book.Year).Scan(&bookID)
	logFatal(err)

	json.NewEncoder(w).Encode(bookID)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello World!")
}
