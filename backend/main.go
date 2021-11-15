package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type counterValue struct {
	Num1      int  `json:"num1"`
	Increased bool `json:"inc"`
	Decreased bool `json:"dec"`
}

type valueResponseData struct {
	Add int `json:"add"`
	Sub int `json:"sub"`
}

func process(numsdata counterValue) valueResponseData {

	var numsres valueResponseData
	if numsdata.Increased {
		numsres.Add = numsdata.Num1 + 1
	}
	if numsdata.Decreased {
		numsres.Sub = numsdata.Num1 - 1
	}

	return numsres
}

func calc(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var numsData counterValue
	var numsResData valueResponseData

	decoder.Decode(&numsData)

	numsResData = process(numsData)

	fmt.Println(numsResData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(numsResData); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", calc)
	http.ListenAndServe(":8090", nil)
}
