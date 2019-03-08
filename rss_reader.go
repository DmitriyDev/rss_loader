package main

import "fmt"

type RssReader interface {
	read(map[string]RssStream)
}

type Reader struct {
	config       Config
	communicator Communicator
	storage      Storage
}

func (r *Reader) read() {

	var readScenario RssReader

	if r.config.Async {
		readScenario = r.asyncReader()
	} else {
		readScenario = r.syncReader()
	}

	readScenario.read(r.config.Streams)
}

func (r *Reader) syncReader() *SyncRssReader {

	fmt.Println("Syncroniouse request")

	return &SyncRssReader{
		config:       r.config,
		communicator: r.communicator,
		storage:      r.storage,
	}

}

func (r *Reader) asyncReader() *AsyncRssReader {

	fmt.Println("Asyncroniouse request")

	return &AsyncRssReader{
		config:       r.config,
		communicator: r.communicator,
		storage:      r.storage,
	}

}
