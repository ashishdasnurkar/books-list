package main

import (
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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", r)
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {

	io.WriteString(writer, "Hello World!")
}
