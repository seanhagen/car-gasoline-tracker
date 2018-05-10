package db

import (
	"database/sql/driver"
	"fmt"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
	uuid "github.com/satori/go.uuid"
)

// UUID wrapper for UUID that handles DB returning uuids as string thanks to pgx
type UUID struct {
	uuid.UUID
}

// Scan - implement the database/sql scanner interface
func (u UUID) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return fmt.Errorf("value is not string")
	}

	id, err := uuid.FromString(v)
	if err != nil {
		return err
	}

	u = UUID{id}
	return nil
}

// Value - implementation of valuer for database/sql
func (u UUID) Value() (driver.Value, error) {
	return fmt.Sprintf("%v", u), nil
}

// UUIDArray wrapper for UUID that handles array of uuids being returned from the db thanks to pgx
type UUIDArray []uuid.UUID

// Scan - implement the database/sql scanner interface
func (ua UUIDArray) Scan(value interface{}) error {
	log.Printf("scanning from value: %v#", value)
	v, ok := value.(string)
	if !ok {
		log.Printf("no bad")
		return fmt.Errorf("value is not string")
	}

	v = strings.Replace(v, "{", "", -1)
	v = strings.Replace(v, "}", "", -1)

	bits := strings.Split(v, ",")
	for _, x := range bits {
		id, err := uuid.FromString(x)
		if err != nil {
			log.Printf("error here: %v", err)
			return err
		}
		ua = append(ua, id)
	}
	log.Printf("got value: %#v", ua)
	spew.Dump(ua)
	return nil
}

// Value - implementation of valuer for database/sql
func (ua UUIDArray) Value() (driver.Value, error) {
	out := []string{}
	for _, x := range ua {
		out = append(out, x.String())
	}
	return fmt.Sprintf("{%v}", strings.Join(out, ",")), nil
}
