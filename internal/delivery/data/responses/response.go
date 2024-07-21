package response

import (
	"encoding/json"
	"net/http"
	"news-topic-api/common"
)

type Response struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data,omitempty"`
	Meta    *common.Meta `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResponseSuccess(w http.ResponseWriter, statusCode int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonResponse, _ := json.MarshalIndent(res, "", "  ")
	w.Write(jsonResponse)
}

func NewResponseError(w http.ResponseWriter, statusCode int, errRes *ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonResponse, _ := json.MarshalIndent(errRes, "", "  ")
	w.Write(jsonResponse)
}
