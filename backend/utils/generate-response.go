package utils

import (
	"encoding/json"
	"net/http"
)

type error struct {
	ErrorMessage string `json:"errorMessage"`
}

func CreateErrorResponse(response http.ResponseWriter, statusCode int, errorMessage string) {
	errDetail := error{
		ErrorMessage: errorMessage,
	}
	response.WriteHeader(statusCode)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(errDetail)
}

type success struct {
	SuccessMessage string      `json:"successMessage"`
	Data           interface{} `json:"data"`
}

func CreateSuccessResponse(response http.ResponseWriter, successMessage string, data interface{}) {
	successDetail := success{
		SuccessMessage: successMessage,
		Data:           data,
	}
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(successDetail)
}
