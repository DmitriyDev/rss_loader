package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Target_website string
	Store_folder   string
	Async          bool
	Abs_path       bool
	Streams        map[string]RssStream
}

const CONFIG_FILE = "/config.yml"

var globalConfig Config

func init()  {

}

func getConfig() Config {

	var config Config

	filename := getWorkDir(true) + CONFIG_FILE

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {

		filename := getWorkDir(false) + CONFIG_FILE
		yamlFile, err = ioutil.ReadFile(filename)
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
