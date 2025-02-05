package marshalizers

import (
    "encoding/json"
    "io"
)

func UnmarshalJson[T any](r io.Reader) (*T, error) {
    var target T
    err := json.NewDecoder(r).Decode(&target)
    return &target, err
}

func MarshalJson[T any](w io.Writer, msg T) error {
    return json.NewEncoder(w).Encode(msg)
}
