package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type errResponse struct {
	Error string `json:"error"`
}

func JSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		fmt.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "15")

	w.WriteHeader(code)
	w.Write(data)
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		fmt.Println("Responding with 5xx error", message)
	}

	JSONResponse(w, code, errResponse{
		Error: message,
	})

}

type ReturnTokenResponseStruct struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

func ReturnTokenResponse(token string, user User) ReturnTokenResponseStruct {
	return ReturnTokenResponseStruct{
		Token: token,
		User:  user,
	}
}

type APIQueryResponseStruct struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
	CurrentPage int32 `json:"page"`
	TotalPages  int64 `json:"total_pages"`
	Limit       int64 `json:"limit"`
}
