package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
	default:
		return fmt.Errorf("invalid duration type: %T", value)
	}
	return nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(d.Duration.Seconds()))
}

func (d *Duration) Scan(value interface{}) error {
	switch v := value.(type) {
	case nil:
		d.Duration = 0
	case Duration:
		d.Duration = v.Duration
	case int64:
		d.Duration = time.Duration(v) * time.Second
	default:
		return fmt.Errorf("cannot convert %T to Duration", value)
	}
	return nil
}

func (d Duration) Value() (driver.Value, error) {
	return int64(d.Duration.Seconds()), nil
}
