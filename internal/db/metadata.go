package db

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Metadata TODO
type Metadata map[string]interface{}

// Value TODO
func (m Metadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// Scan TODO
func (m *Metadata) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*m, ok = i.(map[string]interface{})
	if !ok {
		return fmt.Errorf("type assertion .(map[string]interface{}) failed.")
	}

	return nil
}
