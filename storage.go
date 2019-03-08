package main

import (
	"io/ioutil"
	"fmt"
	"path/filepath"
)

type Storage interface {
	save(file string, content string)
}

const PATH_DELIMITER = "/"
const EXTENTION = ".xml"
const PATH_PREFIX = "."

type FileStorage struct {
	config Config
}

func (fs FileStorage) filePath(file string) string {

	conf := fs.config

	path := PATH_PREFIX + PATH_DELIMITER + conf.Store_folder + PATH_DELIMITER + file + EXTENTION

	abs_fpath, _ := filepath.Abs(path)

	return abs_fpath
}

func (fs FileStorage) save(file string, content string) {

	conf := fs.config

	bcontent := []byte(content)

	err := ioutil.WriteFile("./"+conf.Store_folder+"/"+file+".xml", bcontent, 0644)

	if err != nil {
		fmt.Println("Error", err)
	}
}
