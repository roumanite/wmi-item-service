package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config ...
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		DbName   string `yaml:"db_name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func LoadConfig() Config {
  var configuration Config
	f, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(f, &configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}
