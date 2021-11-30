package main

import (
	"counterapi/middlewares"
	"net/http"
)

func main() {

	http.HandleFunc("/", middlewares.Calc)
	http.ListenAndServe(":8090", nil)
}
