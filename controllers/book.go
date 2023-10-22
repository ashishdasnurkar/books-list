package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ashishdasnurkar/books-list/models"
)

type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}

		rows, err := db.Query("select * from books")
		logFatal(err)

		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
			logFatal(err)

			books = append(books, book)
		}

		json.NewEncoder(w).Encode(books)
	}
}
