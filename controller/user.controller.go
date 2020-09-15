package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/young-seung/msa-example/account/account/command"
)

// Controller http controller
type Controller interface{}

// ControllerImplement http controller implement
type ControllerImplement struct{}

func (controller ControllerImplement) handle(responseWriter http.ResponseWriter, request *http.Request) {

}

func branchByMethod(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		handlePOST(responseWriter, request)
	case "PUT":
		handlePUT(responseWriter, request)
	case "DELETE":
		handleDELETE(responseWriter, request)
	case "GET":
		handleGET(responseWriter, request)
	default:
		handleNotAllowedMethod(responseWriter, request)
	}
}

type POSTBody struct {
	Email    string
	Password string
}

func handlePOST(responseWriter http.ResponseWriter, request *http.Request) {
	contentLength := request.ContentLength
	body := POSTBody{}
	decodeJSONBody(responseWriter, request, body)
	command := command.CreateCommand()
}

func handlePUT(responseWriter http.ResponseWriter, request *http.Request) {

}

func handleDELETE(responseWriter http.ResponseWriter, request *http.Request) {

}

func handleGET(responseWriter http.ResponseWriter, request *http.Request) {

}

func handleNotAllowedMethod(responseWriter http.ResponseWriter, request *http.Request) {

}

func createCommand() {

}

func createQuery() {

}

type httpError struct {
	status  int
	message string
}

func (err *httpError) Error() string {
	return err.message
}

func decodeJSONBody(responseWriter http.ResponseWriter, request *http.Request, destination interface{}) error {
	if request.Header.Get("Content-Type") != "" {
		message := "Content-Type header is not application/json"
		return &httpError{status: http.StatusUnsupportedMediaType, message: message}
	}

	request.Body = http.MaxBytesReader(responseWriter, request.Body, 1048576)

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&destination); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			message := strings.Join([]string("Request body contains badly-formed JSON (at position "), syntaxError.Offset)
			return &httpError{status: http.StatusBadRequest, message: message}

		case errors.Is(err, io.ErrUnexpectedEOF):
			message := "Request body contains badly-formed JSON"
			return &httpError{status: http.StatusBadRequest, message: message}

		case errors.As(err, &unmarshalTypeError):
			message := strings.Join([]string("Request body containers an invalid value for the ", unmarshalTypeError.Field, "field (at position )", unmarshalTypeError.Offset))
			return &httpError{status: http.StatusBadRequest, message: message}

		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			message := strings.Join([]string("Request body contains unknown field ", fieldName))
			return &httpError{status: http.StatusBadRequest, message: message}

		case errors.Is(err, io.EOF):
			message := "Request body must not be empty"
			return &httpError{status: http.StatusBadRequest, message: message}

		case err.Error() == "http: request body too large":
			message = "Request body must not be larger than 1MB"
			return &httpError{status: http.StatusRequestEntityTooLarge, message: message}

		default:
			return err
		}

		if err = decoder.Decode(&struct{}{}); io.EOF {
			message := "Request body must only contain a single JSON object"
			return &httpError{status: http.StatusBadRequest, message: message}
		}

		return nil
	}
}
