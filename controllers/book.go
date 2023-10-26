package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ashishdasnurkar/books-list/models"
	bookRepository "github.com/ashishdasnurkar/books-list/repository/book"
	"github.com/ashishdasnurkar/books-list/utils"
	"github.com/gorilla/mux"
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
		var error models.Error
		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			error.Message = "Internal server error ..."
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		params := mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}
		book, err := bookRepo.GetBook(db, book, params["id"])

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "No book found ..."
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Internal server error ..."
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)

	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}
		rowsUpdated, err := bookRepo.RemoveBook(db, params["id"])

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "No book found ..."
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Internal server error ..."
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rowsUpdated)
	}
}
