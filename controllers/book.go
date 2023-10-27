package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

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

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int
		var error models.Error
		json.NewDecoder(r.Body).Decode(&book)

		if strings.TrimSpace(book.Author) == "" || strings.TrimSpace(book.Title) == "" || strings.TrimSpace(book.Year) == "" {
			error.Message = "Missing fields ..."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book, bookID)

		if err != nil {
			error.Message = "Internal server error ..."
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var rowsUpdated int64
		var error models.Error
		json.NewDecoder(r.Body).Decode(&book)

		if book.ID == "0" || strings.TrimSpace(book.Author) == "" || strings.TrimSpace(book.Title) == "" || strings.TrimSpace(book.Year) == "" {
			error.Message = "Missing fields ..."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			error.Message = "Internal server error ..."
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rowsUpdated)
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
