package main

import (
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Target_website string
	Store_folder   string
	Async          bool
	Streams        map[string]RssStream
}

func getConfig() Config {

	config_file := "./config.yml"

	var config Config

	filename, _ := filepath.Abs(config_file)
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
