package main

import (
	"io/ioutil"
	"fmt"
	"path/filepath"
	"os"
)

type Storage interface {
	save(file string, content string)
}

const PATH_DELIMITER = "/"
const EXTENTION = ".xml"
const PERMISSIONS = 0755

type FileStorage struct {
	config Config
}

func (fs FileStorage) checkStoreFolder() bool {
	if _, err := os.Stat(fs.storeFolder()); os.IsNotExist(err) {
		return false
	}

	return true
}

func (fs FileStorage) makeStoreFolder() {
	os.Mkdir(fs.storeFolder(), PERMISSIONS)
}

func (fs FileStorage) storeFolder() string {

	conf := fs.config

	path := getWorkDir(conf.Abs_path) + PATH_DELIMITER + conf.Store_folder

	return path
}

func (fs FileStorage) filePath(file string) string {

	path := fs.storeFolder() + PATH_DELIMITER + file + EXTENTION

	abs_fpath, _ := filepath.Abs(path)

	return abs_fpath
}

func (fs FileStorage) save(file string, content string) {

	if fs.checkStoreFolder() == false {
		fs.makeStoreFolder()
	}

	bcontent := []byte(content)

	file_path := fs.filePath(file)

	err := ioutil.WriteFile(file_path, bcontent, PERMISSIONS)

	if err != nil {
		fmt.Println("Error", err)
	}
}
