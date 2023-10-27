package main

import (
	"database/sql"
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
	r.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	r.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	r.HandleFunc("/book/{id}", controller.RemoveBook(db)).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello to the Books REST endpoint server!")
}
