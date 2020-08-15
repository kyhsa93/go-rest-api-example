package main

import (
	"log"
	"net/http"
	"strings"
)

type controller struct{}

func (c controller) handle(responseWriter http.ResponseWriter, request *http.Request) {

	log.Println(request.Method) // GET, POST, PUT, DELETE

	log.Println(strings.TrimPrefix(request.URL.Path, "/users/")) // "", "...data"

	data := []byte("handled")
	responseWriter.Write(data)
}

func main() {
	controller := controller{}
	http.HandleFunc("/users/", controller.handle)

	http.ListenAndServe(":5000", nil)
}
