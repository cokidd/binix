package api

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"messages"`
	Data    any    `json:"data,omitempty"`
}

func Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit at Broker",
	}
	out, _ := json.MarshalIndent(payload, "", "\t")
	w.WriteHeader(http.StatusAccepted)

	w.Write(out)
}
