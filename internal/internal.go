package internal

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gchaincl/dotsql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // sql driver
)

// Config contains the configuration needed by all the routes
type Config struct {
	HTTP *http.Client
	DB   *sqlx.DB
	Dot  *dotsql.DotSql

	AppID   string
	AppName string
	Env     string
	Port    string
	Project string

	_store map[string]interface{}
}

// NewConfig TODO
func NewConfig(appName string) (*Config, error) {
	e := os.Getenv("GOOGLE_ENV")
	e = strings.TrimSpace(e)
	if e == "" {
		e = "development"
	}

	http := setupHTTPClient(goticFiles)
	port := setupPort()

	c := &Config{
		HTTP:    http,
		AppName: appName,
		Port:    port,
		Env:     e,
		Project: os.Getenv("GOOGLE_PROJECT"),
	}

	db, dot, err := setupDatabase()
	if err != nil {
		return nil, err
	}

	c.DB = db
	c.Dot = dot

	return c, nil
}

// IsProd ...
func (c *Config) IsProd() bool {
	return c.Env == "production"
}

// Shutdown ...
func (c *Config) Shutdown() {
	err := c.DB.Close()
	if err != nil {
		log.Printf("Unable to close datastore connection: %v", err)
	}
}

// Set ...
func (c *Config) Set(name string, in interface{}) {
	c._store[name] = in
}

// Get ...
func (c *Config) Get(name string) interface{} {
	return c._store[name]
}

// setupHTTPClient sets up the HTTP client with the built-in certs
func setupHTTPClient(files map[string]string) *http.Client {
	pool := x509.NewCertPool()
	buf := bytes.NewBufferString(files["config/ca-certificates.crt"])
	pool.AppendCertsFromPEM(buf.Bytes())
	return &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{RootCAs: pool}}}
}

func setupPort() string {
	p := os.Getenv("PORT")
	if p == "" {
		p = "8080"
	}
	return p
}

func setupDatabase() (*sqlx.DB, *dotsql.DotSql, error) {
	dsn := mustGetenv("POSTGRES_CONNECTION")
	log.Printf("got dsn: %v", dsn)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, nil, err
	}

	dot, err := dotsql.LoadFromString(goticFiles["db/final.sql"])
	if err != nil {
		return nil, nil, err
	}

	return db, dot, nil
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
