package response

import (
    "dockMon/pkg/marshalizers"
    "net/http"
)

type Error struct {
    Message string `json:"message"`
}

func InternalServerError(w http.ResponseWriter, message string) {
    ErrorResponse(w, http.StatusInternalServerError, message)
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
    w.WriteHeader(code)
    _ = marshalizers.MarshalJson(w, Error{Message: message})
}
