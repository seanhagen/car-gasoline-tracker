package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gchaincl/dotsql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // sql driver
	"github.com/satori/go.uuid"
	"github.com/seanhagen/gas-web/internal/db"
)

type (
	// Storage is a concrete implementation of the db.Storage interface
	Storage struct {
		dot  *dotsql.DotSql
		db   *sqlx.DB
		last sql.Result
	}
)

// SetupDB sets up the database connection for us
func SetupDB(queries string) (*Storage, error) {
	dsnBase := "postgres://%s:%s@%s:%s/%s?sslmode=disable"
	dsnBlank := fmt.Sprintf(dsnBase, "", "", "", "", "")
	dsn := fmt.Sprintf(
		dsnBase,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	if dsn == dsnBlank {
		return nil, errors.New("unable to connect, one of DB_HOST, DB_PORT, DB_NAME, DB_USER, or DB_PASS was blank")
	}

	pgs := Storage{}
	err := pgs.Connect(dsn)
	if err != nil {
		return nil, err
	}

	d, err := dotsql.LoadFromString(queries)
	if err != nil {
		return nil, err
	}
	pgs.dot = d
	return &pgs, nil
}

func retryConnect(attempts int, sleep time.Duration, callback func() error) (err error) {
	for i := 0; ; i++ {
		err = callback()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)
		log.Println("retrying to connect to the database after error: ", err)
	}
	return fmt.Errorf("after %d attemps, last error: %s", attempts, err)
}

// Connect fulfills Storage interface. It will retry multiple times to connect to the database,
// the number of times it tries can be controlled using the DB_RETRY environment variable. By default,
// it'll retry 20 times.
func (pgs *Storage) Connect(uri string) error {
	retry := os.Getenv("DB_RETRY")

	retryCount := 20
	var err error
	if retry != "" {
		rc, err := strconv.Atoi(retry)
		if err != nil {
			log.Printf("Unable to convert DB_RETRY to integer, got: %s, using default value", retry)
		} else {
			retryCount = rc
		}
	}

	var db *sqlx.DB
	err = retryConnect(retryCount, 2*time.Second, func() (err error) {
		x, err := sqlx.Connect("postgres", uri)
		if err == nil {
			db = x
		}
		return
	})

	if err != nil {
		return err
	}

	db.SetMaxOpenConns(30)

	pgs.db = db

	return nil
}

// Close TODO
func (pgs *Storage) Close() error {
	return pgs.db.Close()
}

// Begin TODO
func (pgs *Storage) Begin() (db.TX, error) {
	tx, err := pgs.db.Beginx()
	if err != nil {
		return nil, err
	}

	x, err := createTransaction(tx, pgs.dot)
	if err != nil {
		return nil, err
	}

	return x, nil
}

// LastResult fulfills the Storage interface
func (pgs *Storage) LastResult() sql.Result {
	return pgs.last
}

// Raw fulfills the Storage interface
func (pgs *Storage) Raw(name string) (string, error) {
	return pgs.dot.Raw(name)
}

// Exec fulfills the Storage interface
func (pgs *Storage) Exec(queryName string, args ...interface{}) error {
	query, err := pgs.dot.Raw(queryName)
	if err != nil {
		return err
	}

	last, err := pgs.db.Exec(query, args...)
	pgs.last = last
	return err
}

// MustExec attempts to execute the named query, and if there are errors it will panic
func (pgs *Storage) MustExec(queryName string, args ...interface{}) {
	query, err := pgs.dot.Raw(queryName)
	if err != nil {
		panic(err)
	}

	last := pgs.db.MustExec(query, args...)
	pgs.last = last
}

// NamedExec fulfills the Storage interface
func (pgs *Storage) NamedExec(queryName string, arg interface{}) error {
	query, err := pgs.dot.Raw(queryName)
	if err != nil {
		return err
	}

	last, err := pgs.db.NamedExec(query, arg)
	pgs.last = last
	return err
}

// NamedQuery TODO
func (pgs *Storage) NamedQuery(queryName string, arg interface{}) (*sqlx.Rows, error) {
	query, err := pgs.dot.Raw(queryName)
	if err != nil {
		return nil, err
	}

	return pgs.db.NamedQuery(query, arg)
}

// Getter fulfils the Storage interface
func (pgs *Storage) Getter(i interface{}, queryName string, args ...interface{}) error {
	query, err := pgs.dot.Raw(queryName)
	if err != nil {
		return err
	}
	return pgs.db.Get(i, query, args...)
}

// Select fulfills the Storage interface
func (pgs *Storage) Select(i interface{}, queryName string, args ...interface{}) error {
	query, err := pgs.dot.Raw(queryName)
	if err != nil {
		return err
	}

	if strings.Contains(query, "(?)") {
		query, args, err = sqlx.In(query, args...)
		if err != nil {
			return err
		}
	}
	query = pgs.db.Rebind(query)

	err = pgs.db.Select(i, query, args...)
	if err != nil {
		log.Printf("postgres select, error: '%v', query: '%v', args:", err, query)
		spew.Dump(args)
	}
	return err
}

// SelectPage fulfills the Storage interface
func (pgs *Storage) SelectPage(results interface{}, fa db.FilterArgs) error {
	q, err := fa.Query()
	if err != nil {
		return err
	}

	query, args, err := pgs.fixQuery(q, fa.Args())
	if err != nil {
		return err
	}

	err = pgs.db.Select(results, query, args...)
	if err != nil {
		log.Printf("postgres select, error: '%v', query: '%v', args:", err, query)
		spew.Dump(args)
	}
	return err
}

// PageCount TODO
func (pgs *Storage) PageCount(fa db.FilterArgs) (int, error) {
	q, err := fa.Count()
	if err != nil {
		return 0, err
	}

	countQuery, args, err := pgs.fixQuery(q, fa.Args())
	if err != nil {
		return 0, err
	}

	var count int
	err = pgs.db.Get(&count, countQuery, args...)
	return count, err
}

// fixQuery TODO
func (pgs *Storage) fixQuery(query string, x map[string]interface{}) (string, []interface{}, error) {
	query, args, err := sqlx.Named(query, x)
	if err != nil {
		return "", []interface{}{}, err
	}

	if strings.Contains(query, "(?)") {
		query, args, err = sqlx.In(query, args...)
		if err != nil {
			return "", []interface{}{}, err
		}
	}

	query = pgs.db.Rebind(query)
	return query, args, nil
}

// Get fulfills Storage method
func (pgs *Storage) Get(s db.Storable, ids ...uuid.UUID) (db.Storable, error) {
	s.ID(ids...)
	err := s.Fetch(pgs)
	return s, err
}

// Insert fulfills Storage method
func (pgs *Storage) Insert(s db.Storable) error {
	return s.Insert(pgs)
}

// Update fulfills Storage method
func (pgs *Storage) Update(s db.Storable) error {
	return s.Update(pgs)
}

// Delete fulfills Storage method
func (pgs *Storage) Delete(s db.Storable) error {
	return s.Delete(pgs)
}
