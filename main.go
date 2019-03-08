package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Start process")

	c := getConfig()

	reader := RssReaderStrategy{
		config:       c,
		communicator: DZoneCommunicator{c},
	}

	streamContents := reader.contentOfStreams()

	storage := FileStorage{c}

	for _, streamContent := range streamContents {
		storage.save(streamContent.stream.Code, streamContent.content)
	}

	fmt.Println(time.Since(start).Seconds())

}

