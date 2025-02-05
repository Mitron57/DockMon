package models

import (
    "database/sql"
    "encoding/json"
    "time"
)

type Time struct {
    sql.NullTime
}

func (t Time) MarshalJSON() ([]byte, error) {
    if t.Valid {
        return json.Marshal(t.Time)
    }
    return json.Marshal(nil)
}

func (t *Time) UnmarshalJSON(data []byte) error {
    var v time.Time
    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }
    t.NullTime = sql.NullTime{Time: v, Valid: true}
    return nil
}
