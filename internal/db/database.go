package db

import (
	"net/url"

	"github.com/satori/go.uuid"
)

type (
	// FilterArgs provides arguments to queries if required
	FilterArgs interface {
		Populate(url.Values) ErrorCollector
		// Query returns two queries -- one with limits for pagination, one without for counting
		Query() (string, error)
		Count() (string, error)
		Args() map[string]interface{}
	}

	// Storable is an interface for models that are going to be created/updated from the data store
	// via the Storage interface.
	Storable interface {
		ID(...uuid.UUID)
		Fetch(ModelFuncs) error
		Insert(ModelFuncs) error
		Update(ModelFuncs) error
		Delete(ModelFuncs) error
	}

	// ModelFuncs is a subset of the methods provided by both TX & Storage, to allow Storable to function with either
	ModelFuncs interface {
		Getter(dest interface{}, queryName string, args ...interface{}) error
		NamedExec(queryName string, arg interface{}) error
		// NamedQuery(queryName string, arg interface{}) (*sqlx.Rows, error)
		Select(dest interface{}, query string, args ...interface{}) error
	}

	// TX defines an interface for setting up a transaction
	TX interface {
		// Getter fetches a single record into the `dest` interface
		Getter(dest interface{}, queryName string, args ...interface{}) error
		// Select fetches at least one record into the interface, expects an array
		Select(dest interface{}, query string, args ...interface{}) error

		// Exec executes the given query with the provided arguments, within the transaction
		Exec(queryName string, args ...interface{}) error

		// NamedExec is basically a wrapper for sqlx.DB.NamedExec, within the transaction
		NamedExec(queryName string, arg interface{}) error

		// // NamedQuery is basically a wrapper for sqlx.DB.NamedQuery, within the transaction
		// NamedQuery(queryName string, arg interface{}) (*sqlx.Rows, error)

		// Fetch finds a Storable, within the transaction
		Fetch(Storable) error

		// Inserts the Storable into the database, within the transaction
		Insert(Storable) error

		// Update updates the Storable in the database, within the transaction
		Update(Storable) error

		// Delete removes the record associated with the Storable, within the transaction
		Delete(Storable) error

		// Commit applies everything that happened within the transaction to the database
		Commit() error

		// Rollback reverts all the changes that happened within the transaction
		Rollback() error
	}

	// Storage is an interface for hiding the nitty-gritty of the database stuff.
	Storage interface {
		// Close shuts down the database interface
		Close() error

		// Begin starts a database transaction (if supported by the underlying data store)
		Begin() (TX, error)

		// // LastResult returns the results of the previous SQL call
		// LastResult() interface{}

		// Connect tells the Storage implementation to attempt to connect to the database
		// and return an error if unable to connect.
		Connect(string) error

		// Returns the raw SQL for the named query
		Raw(string) (string, error)

		// Exec executes the given query with the provided arguments
		Exec(queryName string, args ...interface{}) error

		// MustExec is similar to Exec, except if there are errors it will panic
		MustExec(queryName string, args ...interface{})

		// NamedExec is basically a wrapper for sqlx.DB.NamedExec
		NamedExec(queryName string, args interface{}) error

		// // NamedQuery is a wrapper for sqlx.DB.NamedQuery
		// NamedQuery(queryName string, args interface{}) (*sqlx.Rows, error)

		// Getter is basically a wrapper for sqlx.DB.Get
		Getter(dest interface{}, queryName string, args ...interface{}) error

		// Selecter is basically a wrapper for sqlx.DB.Select
		Select(i interface{}, queryName string, args ...interface{}) error

		// SelectPage loads a page of results from the storage back-end by calling SelectPage on the Storable
		SelectPage(i interface{}, args FilterArgs) error

		// PageCount gets the total count, ignoring pagination limits
		PageCount(args FilterArgs) (int, error)

		// Get loads the Storable from the storage back-end by calling Storable.Fetch on it
		Get(Storable, ...uuid.UUID) (Storable, error)

		// Insert stores the Storable in the storage back-end by calling Insert on the Storable
		Insert(Storable) error

		// Update updates the record in the storabe back-end by calling Update on the Storable
		Update(Storable) error

		// Delete removes the record associated with the Storable by calling Delete on it
		Delete(Storable) error
	}
)
