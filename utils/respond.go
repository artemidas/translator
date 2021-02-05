package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func DecodeBody(r *http.Request, data interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(data)
}

func EncodeBody(w http.ResponseWriter, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func RespondJson(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	EncodeBody(w, data)
}

func RespondHttpError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	EncodeBody(w, &Response{Code: code, Message: message})
}
