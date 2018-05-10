package models

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func Test_VisitBasic(t *testing.T) {
	v := Visit{}

	if v.ID != uuid.Nil {
		t.Error("ID doesn't default to uuid.Nil")
	}
}
