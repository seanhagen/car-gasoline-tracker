package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type DbConfig struct {
	Adapter  string `yaml:"adapter,omitempty"`
	Pool     int
	Timeout  int
	Username string
	Password string
	Database string
}

type GoogleApiConfig struct {
	Key string
}

var dbconfig *DbConfig
var apiconfig *GoogleApiConfig

func loadConfigFromYAML(file string) []byte {
	filename, _ := filepath.Abs(file)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return yamlFile
}

func loadDbConfig() *DbConfig {
	if dbconfig != nil {
		return dbconfig
	}
	bytes := loadConfigFromYAML("./config/db.yml")
	err := yaml.Unmarshal(bytes, &dbconfig)
	if err != nil {
		panic(err)
	}
	return dbconfig
}

func loadGoogleApiConfig() *GoogleApiConfig {
	if apiconfig != nil {
		return apiconfig
	}
	bytes := loadConfigFromYAML("./config/api.yml")
	err := yaml.Unmarshal(bytes, &apiconfig)
	if err != nil {
		panic(err)
	}
	return apiconfig
}

func loadDBConnectionString() string {
	config := loadDbConfig()
	var buffer bytes.Buffer

	buffer.WriteString("user=")
	buffer.WriteString(config.Username)
	buffer.WriteString(" password=")
	buffer.WriteString(config.Password)
	buffer.WriteString(" dbname=")
	buffer.WriteString(config.Database)

	return buffer.String()
}
