package main

import "net/http"

type Controller interface{}

type ControllerImplement struct{}

func (controller ControllerImplement) handle(responseWriter http.ResponseWriter, request *http.Request) {

}

func branchByMethod(request *http.Request) {

}

func createCommand() {

}

func createQuery() {

}
