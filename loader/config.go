package loader

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

type Config struct {
	Target_website string
	Store_folder   string
	Async          bool
	Abs_path       bool
	Streams        map[string]RssStream
}

var globalConfig Config

func getConfig(configFile string) Config {

	var config Config

	filename := filepath.Dir(".") + configFile

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {

		if err != nil {
			panic(err)
		}
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
