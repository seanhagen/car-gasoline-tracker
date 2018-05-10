package datastore

import (
	"cloud.google.com/go/datastore"
	"github.com/gchaincl/dotsql"
	"github.com/jmoiron/sqlx"
	"github.com/seanhagen/gas-web/internal/db"
)

type TX struct {
	x          *datastore.Transaction
	dot        *dotsql.DotSql
	lastCommit *datastore.Commit
}

func createTransaction(ds *Storage) (*TX, error) {
	tx, err := ds.db.NewTransaction(ds.ctx)
	if err != nil {
		return nil, err
	}

	return &TX{x: tx, dot: ds.dot}, nil
}

// Getter TODO
func (t *TX) Getter(dest interface{}, queryName string, args ...interface{}) error {
	return nil
}

// Select TODO
func (t *TX) Select(dest interface{}, queryName string, args ...interface{}) error {
	return nil
}

// Exec TODO
func (t *TX) Exec(queryName string, args ...interface{}) error {
	return nil
}

// NamedExec TODO
func (t *TX) NamedExec(queryName string, arg interface{}) error {
	return nil
}

// NamedQuery TODO
func (t *TX) NamedQuery(queryName string, arg interface{}) (*sqlx.Rows, error) {
	return nil, nil
}

// Fetch TODO
func (t *TX) Fetch(f db.Storable) error {
	return nil
}

// Insert TODO
func (t *TX) Insert(f db.Storable) error {
	return nil
}

// Update TODO
func (t *TX) Update(f db.Storable) error {
	return nil
}

// Delete TODO
func (t *TX) Delete(f db.Storable) error {
	return nil
}

// Commit TODO
func (t *TX) Commit() error {
	c, er := t.x.Commit()
	t.lastCommit = c
	return er
}

// Rollback TODO
func (t *TX) Rollback() error {
	return t.x.Rollback()
}
