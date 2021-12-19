package main

import (
	"counterapi/service"
	"net/http"
)

func main() {

	http.HandleFunc("/", service.Calc)
	http.ListenAndServe(":8090", nil)
}
