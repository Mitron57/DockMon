package response

import (
    "dockMon/pkg/marshalizers"
    "net/http"
)

func WriteJSON(w http.ResponseWriter, code int, payload any) {
    w.WriteHeader(code)
    w.Header().Set("Content-Type", "application/json")
    _ = marshalizers.MarshalJson(w, payload)
}
