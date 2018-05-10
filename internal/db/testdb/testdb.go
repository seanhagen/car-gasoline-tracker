package testdb

import (
	"database/sql"

	"github.com/seanhagen/gas-web/internal/db"
	// "github.com/jmoiron/sqlx/reflectx"
	"log"
	// "reflect"
)

type (
	result struct{}

	// TestDB is a BoltDB-backed store for use in unit tests
	TestDB struct {
		Page interface{}
	}
)

// LastInsertId TODO
func (r result) LastInsertId() (int64, error) {
	return 0, nil
}

// RowsAffected TODO
func (r result) RowsAffected() (int64, error) {
	return 0, nil
}

// LastResult ...
func (db TestDB) LastResult() sql.Result {
	return result{}
}

// Connect TODO
func (db TestDB) Connect(driver, uri string) error {
	return nil
}

// Raw TODO
func (db TestDB) Raw(queryName string) (string, error) {
	return "", nil
}

// Exec TODO
func (db TestDB) Exec(queryName string, args ...interface{}) error {
	return nil
}

// MustExec TODO
func (db TestDB) MustExec(queryName string, args ...interface{}) {
	return
}

// NamedExec TODO
func (db TestDB) NamedExec(queryName string, arg interface{}) error {
	return nil
}

// Getter TODO
func (db TestDB) Getter(i interface{}, queryName string, args ...interface{}) error {
	return nil
}

// Selecter TODO
func (db TestDB) Selecter(i interface{}, queryName string, args ...interface{}) error {
	return nil
}

// SelectPage TODO
func (db TestDB) SelectPage(dest interface{}, fa db.FilterArgs) error {
	log.Printf("dest: %#v\n%#v\n.", &dest, dest)
	log.Printf("page: %#v\n%#v\n.", &db.Page, db.Page)
	dest = db.Page
	log.Printf("dest: %#v\n%#v\n.", &dest, dest)

	return nil
}

// Get TODO
func (db TestDB) Get(s db.Storable) error {
	return nil
}

// Insert TODO
func (db TestDB) Insert(s db.Storable) error {
	return nil
}

// Update TODO
func (db TestDB) Update(s db.Storable) error {
	return nil
}

// Delete TODO
func (db TestDB) Delete(s db.Storable) error {
	return nil
}
