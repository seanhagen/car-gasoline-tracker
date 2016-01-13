package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Adapter  string `yaml:"adapter,omitempty"`
	Pool     int
	Timeout  int
	Username string
	Password string
	Database string
}

var config *Config

func loadConfig() *Config {
	if config != nil {
		return config
	}

	filename, _ := filepath.Abs("./test.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func loadDBConnectionString() string {
	config := loadConfig()
	var buffer bytes.Buffer

	buffer.WriteString("user=")
	buffer.WriteString(config.Username)
	buffer.WriteString(" password=")
	buffer.WriteString(config.Password)
	buffer.WriteString(" dbname=")
	buffer.WriteString(config.Database)

	return buffer.String()
}
