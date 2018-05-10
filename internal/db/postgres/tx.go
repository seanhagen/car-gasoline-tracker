package postgres

import (
	"github.com/gchaincl/dotsql"
	"github.com/jmoiron/sqlx"
	"github.com/seanhagen/gas-web/internal/db"
)

type pgTX struct {
	x   *sqlx.Tx
	dot *dotsql.DotSql
}

func createTransaction(x *sqlx.Tx, d *dotsql.DotSql) (*pgTX, error) {
	return &pgTX{x: x, dot: d}, nil
}

// Getter TODO
func (t *pgTX) Getter(dest interface{}, queryName string, args ...interface{}) error {
	query, err := t.dot.Raw(queryName)
	if err != nil {
		return err
	}

	// log.Printf("tx, query: %v", query)
	// spew.Dump(args)

	return t.x.Get(dest, query, args...)
}

// Select TODO
func (t *pgTX) Select(dest interface{}, queryName string, args ...interface{}) error {
	query, err := t.dot.Raw(queryName)
	if err != nil {
		return err
	}
	return t.x.Select(dest, query, args...)
}

// Exec TODO
func (t *pgTX) Exec(queryName string, args ...interface{}) error {
	query, err := t.dot.Raw(queryName)
	if err != nil {
		return err
	}
	_, err = t.x.Exec(query, args...)
	return err
}

// NamedExec TODO
func (t *pgTX) NamedExec(queryName string, arg interface{}) error {
	query, err := t.dot.Raw(queryName)
	if err != nil {
		return err
	}
	_, err = t.x.NamedExec(query, arg)
	return err
}

// // NamedQuery TODO
// func (t *pgTX) NamedQuery(queryName string, arg interface{}) (*sqlx.Rows, error) {
// 	query, err := t.dot.Raw(queryName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return t.x.NamedQuery(query, arg)
// }

// Fetch TODO
func (t *pgTX) Fetch(s db.Storable) error {
	return s.Fetch(t)
}

// Insert TODO
func (t *pgTX) Insert(s db.Storable) error {
	return s.Insert(t)
}

// Update TODO
func (t *pgTX) Update(s db.Storable) error {
	return s.Update(t)
}

// Delete TODO
func (t *pgTX) Delete(s db.Storable) error {
	return s.Delete(t)
}

// Commit TODO
func (t *pgTX) Commit() error {
	return t.x.Commit()
}

// Rollback TODO
func (t *pgTX) Rollback() error {
	return t.x.Rollback()
}
