package json

import (
	"encoding/json"
	"fmt"
	"time"
)

// Date TODO
type Date time.Time

// MarshalJSON TODO
func (t Date) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%#v", time.Time(t).Format("2006-01-02"))
	return []byte(stamp), nil
}

// UnmarshalJSON TODO
func (t *Date) UnmarshalJSON(arg []byte) error {
	var s string
	err := json.Unmarshal(arg, &s)
	if err != nil {
		return err
	}
	x, err := time.Parse("2006-01-02", s)
	// log.Printf("json.date - unmarshall, %v, %v, %v, %v", arg, s, x, x.Format("2006-01-02"))
	*t = Date(x)
	return err
}
