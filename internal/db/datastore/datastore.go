package datastore

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gchaincl/dotsql"
	"github.com/satori/go.uuid"
	"github.com/seanhagen/gas-web/internal/db"
)

type (
	// Storage is a concrete implementation of the db.Storage interface
	// specifically for Google Cloud DataStore
	Storage struct {
		dot *dotsql.DotSql
		db  *datastore.Client
		ctx context.Context
	}
)

// SetupDatastore creates a connection to the Google Cloud DataStore
func SetupDatastore(queries, project string) (*Storage, error) {
	project = strings.TrimSpace(project)
	if project == "" {
		return nil, errors.New("unable to connect, GOOGLE_PROJECT must be set and not be blank")
	}

	st := Storage{}
	err := st.Connect(project)
	if err != nil {
		return nil, err
	}

	d, err := dotsql.LoadFromString(queries)
	if err != nil {
		return nil, err
	}

	st.dot = d
	return &st, nil
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
// the number of times it tries can be controlled by using the DB_RETRY environment variable.
//
// By default, it'll attempt to connect 5 times.
func (ds *Storage) Connect(project string) error {
	retry := os.Getenv("DB_RETRY")

	retryCount := 5
	var err error
	if retry != "" {
		rc, err := strconv.Atoi(retry)
		if err != nil {
			log.Printf("Unable to convert DB_RETRY to integer, got %#v, using default value: %v", retry, retryCount)
		} else {
			retryCount = rc
		}
	}

	var client *datastore.Client
	ctx := context.Background()
	err = retryConnect(retryCount, 2*time.Second, func() (err error) {
		dsClient, err := datastore.NewClient(ctx, project)
		if err == nil {
			client = dsClient
		}

		t, err := client.NewTransaction(ctx)
		if err != nil {
			return fmt.Errorf("Google Cloud DataStore: could not connect: %v", err)
		}
		if err := t.Rollback(); err != nil {
			return fmt.Errorf("Google Cloud DataStore: could not connect: %v", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	// log.Printf("got datastore connection")
	// spew.Dump(client)

	ds.db = client
	ds.ctx = ctx

	return nil
}

// Close fulfils Storage interface. Used to shut down datastore connection.
func (ds *Storage) Close() error {
	return ds.db.Close()
}

// Begin fulfils the Storage interface.
func (ds *Storage) Begin() (db.TX, error) {
	x, err := createTransaction(ds)
	if err != nil {
		return nil, err
	}

	return x, nil
}

// Raw TODO
func (ds *Storage) Raw(name string) (string, error) {
	return ds.dot.Raw(name)
}

// Exec TODO
func (ds *Storage) Exec(queryName string, args ...interface{}) error {
	return fmt.Errorf("Function not implemented")
}

// MustExec TODO
func (ds *Storage) MustExec(queryName string, args ...interface{}) {}

// NamedExec TODO
func (ds *Storage) NamedExec(queryName string, arg interface{}) error {
	return fmt.Errorf("Function not implemented")
}

// Getter TODO
func (ds *Storage) Getter(dest interface{}, queryName string, args ...interface{}) error {
	return fmt.Errorf("Function not implemented")
}

// Select TODO
func (ds *Storage) Select(i interface{}, queryName string, args ...interface{}) error {
	return fmt.Errorf("Function not implemented")
}

// SelectPage TODO
func (ds *Storage) SelectPage(i interface{}, args db.FilterArgs) error {
	return fmt.Errorf("Function not implemented")
}

// PageCount TODO
func (ds *Storage) PageCount(args db.FilterArgs) (int, error) {
	return 0, nil
}

// Get TODO
func (ds *Storage) Get(g db.Storable, ids ...uuid.UUID) (db.Storable, error) {
	return nil, fmt.Errorf("Function not implemented")
}

// Insert TODO
func (ds *Storage) Insert(i db.Storable) error {
	return fmt.Errorf("Function not implemented")
}

// Update TODO
func (ds *Storage) Update(u db.Storable) error {
	return fmt.Errorf("Function not implemented")
}

// Delete TODO
func (ds *Storage) Delete(d db.Storable) error {
	return fmt.Errorf("Function not implemented")
}
