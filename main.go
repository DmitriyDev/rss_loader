package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Start process")

	config := getConfig()

	if config.Async {
		processAsync(config)
	} else {
		processSync(config)

	}

	fmt.Println(time.Since(start).Seconds())

}

func processSync(c Config) {

	fmt.Println("Syncroniouse request")

	fs := FileStorage{c}
	communicator := DZoneCommunicator{c}

	rssReader := SyncRssReader{
		config:       c,
		communicator: communicator,
		storage:      fs,
	}

	rssReader.read(c.Streams)
}

func processAsync(c Config) {

	fmt.Println("Asyncroniouse request")

	fs := FileStorage{c}
	communicator := DZoneCommunicator{c}

	rssReader := AsyncRssReader{
		config:       c,
		communicator: communicator,
		storage:      fs,
	}
	rssReader.read(c.Streams)
}
