package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Configuration struct for application
type Configuration struct {
	Server   *Server   `yaml:"server,omitempty"`
	Database *Database `yaml:"database,omitempty"`
}

//Server configuration struct
type Server struct {
	Port string `yaml:"port,omitempty"`
}

//Database configuration struct
type Database struct {
	Type string `yaml:"type,omitempty"`
	PSN  string `yaml:"psn,omitempty"`
}

//Load configuration based on *.yaml file and send path for that file
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	cfg = &Configuration{
		Database: &Database{
			Type: "mysql",
		},
	}
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return cfg, nil
}
