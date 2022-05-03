package http

import (
	"encoding/json"
	"net/http"
)

const (
	JSONContentType = "application/json"
	DateLayout      = "2006-01-02T15:04:05Z"
)

type Error struct {
	Reason string `json:"reason"`
}

type AuthContextKey string

var ContextAccountID AuthContextKey = "account_id"

func SendResponse(w http.ResponseWriter, responseBody interface{}, statusCode int) error {
	w.Header().Set("Content-Type", JSONContentType)
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(responseBody)
}
