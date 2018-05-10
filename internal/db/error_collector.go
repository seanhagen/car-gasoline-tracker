package db

import (
	"fmt"
)

// ErrorCollector TODO
type ErrorCollector []error

func (ce *ErrorCollector) Collect(e error) {
	*ce = append(*ce, e)
}

func (ce ErrorCollector) Error() string {
	err := "Multiple errors:\n"
	for i, e := range ce {
		err += fmt.Sprintf("\tError %d: %s\n", i, e.Error())
	}
	return err
}
