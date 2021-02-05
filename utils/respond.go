package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeBody(r *http.Request, data interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(data)
}
