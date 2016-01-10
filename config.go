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

func loadDBConnectionString() (string) {
	filename, _ := filepath.Abs("./test.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer

	buffer.WriteString("user=")
	buffer.WriteString(config.Username)
	buffer.WriteString(" password=")
	buffer.WriteString(config.Password)
	buffer.WriteString(" dbname=")
	buffer.WriteString(config.Database)

	return buffer.String()
}
