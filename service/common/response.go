package common

import (
	"encoding/json"
	"net/http"
	"time"
)

type Head struct {
	Timestamp time.Time `json:"timestamp"`
}

// such response structure not necessary here, but might be needed for future extension
type Response struct {
	Head *Head       `json:"head"`
	Body interface{} `json:"body"`
}

func (r *Response) WriteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(r)
}

func NewResponse(body interface{}) *Response {
	return &Response{
		Head: &Head{
			Timestamp: time.Now().UTC(),
		},
		Body: body,
	}
}
