package service

import (
	"counterapi/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func process(numsdata models.CounterValue) models.ValueResponseData {

	var numsres models.ValueResponseData
	if numsdata.Increased {
		numsres.Add = numsdata.Num1 + 1
	}
	if numsdata.Decreased {
		numsres.Sub = numsdata.Num1 - 1
	}

	return numsres
}

func Calc(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var numsData models.CounterValue
	var numsResData models.ValueResponseData
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
