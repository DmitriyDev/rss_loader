package rss_loader

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



var globalConfig Config

func init()  {

}

func getConfig(configFile string) Config {

	var config Config

	filename := getWorkDir(true) + configFile

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {

		filename := getWorkDir(false) + configFile
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
