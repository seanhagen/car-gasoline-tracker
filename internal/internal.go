package internal

import (
	"log"

	"github.com/seanhagen/gas-web/internal/app"
)

// Config contains the configuration needed by all the routes
type Config struct {
	app.BaseConfig
}

// NewConfig TODO
func NewConfig(appName string) (*Config, error) {
	bc, err := app.NewBaseConfig(appName, "datastore", goticFiles)
	if err != nil {
		return nil, err
	}
	c := &Config{bc}

	log.Printf("config app id: %v", c.AppID())

	return c, nil
}
