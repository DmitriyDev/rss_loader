package main

import (
	"fmt"
	"time"
	"path/filepath"
	"os"
)

func init() {
	fmt.Println("Start process")
	globalConfig = getConfig()
}

func main() {
	start := time.Now()

	reader := RssReaderStrategy{
		communicator: DZoneCommunicator{globalConfig},
	}

	streamContents := reader.contentOfStreams()

	storage := FileStorage{globalConfig}

	for _, streamContent := range streamContents {
		storage.save(streamContent.stream.Code, streamContent.content)
	}

	fmt.Println(time.Since(start).Seconds())

}

func getWorkDir(absPath bool) string {

	if absPath {
		return filepath.Dir(os.Args[0])
	}

	return filepath.Dir(".")

}
