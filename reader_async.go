package main

import "fmt"

type AsyncRssReader struct {
	config       Config
	communicator Communicator
	storage      Storage
}

func (arr *AsyncRssReader) read(streams map[string]RssStream) {

	ch := make(chan RssStream)

	for _, stream := range streams {
		go arr.asyncRead(stream, ch)
	}

	for _ = range streams {
		stream := <-ch
		fmt.Println(stream.Name + "\t\t\t.... Done")
	}

}

func (arr *AsyncRssReader) asyncRead(stream RssStream, ch chan RssStream) {

	streamUrl := stream.getUrl(arr.config)
	content := arr.communicator.content(streamUrl)
	go arr.storage.save(stream.Code, content)
	ch <- stream

}
