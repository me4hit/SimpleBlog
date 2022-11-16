package main

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		io.WriteString(writer, `{"alive": true}`)
	})
	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", r)
}
