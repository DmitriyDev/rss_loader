package main

import "fmt"

type RssReaderStrategy struct {
	config       Config
	communicator Communicator
}

func (r *RssReaderStrategy) contentOfStreams() []StreamContent {

	var readScenario RssReader

	if r.config.Async {
		readScenario = r.asyncReader()
	} else {
		readScenario = r.syncReader()
	}

	return readScenario.contentOfStreams(r.config.Streams)
}

func (r *RssReaderStrategy) syncReader() *SyncRssReader {

	fmt.Println("Syncroniouse request")

	return &SyncRssReader{
		config:       r.config,
		communicator: r.communicator,
	}

}

func (r *RssReaderStrategy) asyncReader() *AsyncRssReader {

	fmt.Println("Asyncroniouse request")

	return &AsyncRssReader{
		config:       r.config,
		communicator: r.communicator,
	}

}
