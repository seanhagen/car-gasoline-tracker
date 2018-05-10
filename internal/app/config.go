package app

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/seanhagen/gas-web/internal/db"
	"github.com/seanhagen/gas-web/internal/db/datastore"
	"github.com/seanhagen/gas-web/internal/db/postgres"
	"github.com/vmihailenco/msgpack"
)

// IConfig TODO
type IConfig interface {
	Store() db.Storage
	HTTP() *http.Client

	Set(string, interface{})
	Get(string) interface{}

	CacheGet(string, interface{}) error
	CacheSet(string, interface{}, time.Duration) error

	AppID() string
	AppName() string
	Port() string
	IsProd() bool
	Env() string

	Shutdown()
}

// BaseConfig
type BaseConfig struct {
	store db.Storage
	http  *http.Client

	redis *redis.Client
	cache *cache.Codec

	appID   string
	appName string
	env     string
	port    string
	project string

	_store map[string]interface{}
}

// NewBaseConfig TODO
func NewBaseConfig(appName string, dbType string, files map[string]string) (BaseConfig, error) {
	c := BaseConfig{
		_store:  map[string]interface{}{},
		port:    port(),
		project: os.Getenv("GOOGLE_PROJECT"),
	}

	e := os.Getenv("GOOGLE_ENV")
	e = strings.TrimSpace(e)
	if e == "" {
		e = "development"
	}
	c.env = e

	if !c.IsProd() {
		appName = fmt.Sprintf("%v-%v", appName, c.env)
	}
	c.appName = appName

	var store db.Storage
	var err error

	switch dbType {
	case "postgres":
		store, err = setupPostgresDB(files)
	case "datastore":
		store, err = setupDataStoreDB(c.project, files)
	}
	if err != nil {
		return c, fmt.Errorf("Unable to create database connection: %v", err)
	}

	c.store = store
	c.http = setupHTTPClient(files)
	c.redis, c.cache = setupCache()

	return c, nil
}

// Shutdown TODO
func (bc BaseConfig) Shutdown() {
	err := bc.store.Close()
	if err != nil {
		log.Printf("error closing database connection: %v", err)
	}
	err = bc.redis.Close()
	if err != nil {
		log.Printf("error closing redis connectino: %v", err)
	}
}

// AppID TODO
func (bc BaseConfig) AppID() string {
	return bc.appID
}

// AppName TODO
func (bc BaseConfig) AppName() string {
	return bc.appName
}

// Port TODO
func (bc BaseConfig) Port() string {
	return bc.port
}

// Env TODO
func (bc BaseConfig) Env() string {
	return bc.env
}

// IsProd TODO
func (bc BaseConfig) IsProd() bool {
	return bc.env == "production"
}

// Store TODO
func (bc BaseConfig) Store() db.Storage {
	return bc.store
}

// HTTP TODO
func (bc BaseConfig) HTTP() *http.Client {
	return bc.http
}

// Set TODO
func (bc BaseConfig) Set(name string, in interface{}) {
	bc._store[name] = in
}

// Get TODO
func (bc BaseConfig) Get(name string) interface{} {
	return bc._store[name]
}

// CacheGet cache interface TODO
func (bc BaseConfig) CacheGet(key string, object interface{}) error {
	return bc.cache.Get(key, object)
}

// CacheSetWithExpiry TODO
func (bc BaseConfig) CacheSet(key string, object interface{}, expiry time.Duration) error {
	return bc.cache.Set(&cache.Item{
		Key:        key,
		Object:     object,
		Expiration: expiry,
	})
}

// setupCache ...
func setupCache() (*redis.Client, *cache.Codec) {
	r := redis.NewClient(&redis.Options{
		Addr: "redis-master:6379",
	})

	return r, &cache.Codec{
		Redis: r,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},

		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
}

// setDB sets up the PostGres database client
func setupPostgresDB(files map[string]string) (*postgres.Storage, error) {
	store, err := postgres.SetupDB(files["db/final.sql"])
	if err != nil {
		return nil, err
	}
	return store, nil
}

func setupDataStoreDB(project string, files map[string]string) (*datastore.Storage, error) {
	store, err := datastore.SetupDatastore(files["db/final.sql"], project)
	if err != nil {
		return nil, err
	}
	return store, nil
}

// setupHTTPClient sets up the HTTP client with the built-in certs
func setupHTTPClient(files map[string]string) *http.Client {
	pool := x509.NewCertPool()
	buf := bytes.NewBufferString(files["config/ca-certificates.crt"])
	pool.AppendCertsFromPEM(buf.Bytes())
	return &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{RootCAs: pool}}}
}

func port() string {
	p := os.Getenv("PORT")
	if p == "" {
		p = "8080"
	}
	return p
}
