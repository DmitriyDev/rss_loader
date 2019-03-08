package main

import "fmt"

type SyncRssReader struct {
	config       Config
	communicator Communicator
	storage      Storage
}

func (srr *SyncRssReader) read(streams map[string]RssStream) {

	for _, stream := range streams {
		content := srr.syncRead(stream)
		srr.storage.save(stream.Code, content)
		fmt.Println(stream.Name + "\t\t\t.... Done")
	}
}

func (srr *SyncRssReader) syncRead(stream RssStream) string {

	streamUrl := stream.getUrl(srr.config)
	content := srr.communicator.content(streamUrl)
	return content
}
