package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody[T any](r *http.Request) (T, error) {
	var body T
	err := json.NewDecoder(r.Body).Decode(&body)
	return body, err
}
