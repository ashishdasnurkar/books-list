package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ashishdasnurkar/books-list/models"
)

func sendError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func sendSuccess(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}
