package responseBuilder

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(w http.ResponseWriter, message string, data interface{}) {
	write(w, http.StatusOK, message, data)
}

func Created(w http.ResponseWriter, message string, data interface{}) {
	write(w, http.StatusCreated, message, data)
}

func BadRequest(w http.ResponseWriter, message string) {
	write(w, http.StatusBadRequest, message, nil)
}

func Unauthorized(w http.ResponseWriter, message string) {
	write(w, http.StatusUnauthorized, message, nil)
}

func Forbidden(w http.ResponseWriter, message string) {
	write(w, http.StatusForbidden, message, nil)
}

func NotFound(w http.ResponseWriter, message string) {
	write(w, http.StatusNotFound, message, nil)
}

func InternalError(w http.ResponseWriter) {
	write(w, http.StatusInternalServerError, "internal server error", nil)
}

func write(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
