package main

type AsyncRssReader struct {
	config       Config
	communicator Communicator
}

func (arr *AsyncRssReader) contentOfStreams(streams map[string]RssStream) []StreamContent {

	var steamContents []StreamContent
	ch := make(chan StreamContent)

	for _, stream := range streams {
		go arr.asyncRead(stream, ch)
	}

	for _ = range streams {
		streamContent := <-ch
		steamContents = append(steamContents, streamContent)
	}

	return steamContents
}

func (arr *AsyncRssReader) asyncRead(stream RssStream, ch chan StreamContent) {

	streamUrl := stream.getUrl(arr.config)
	content := arr.communicator.content(streamUrl)
	ch <- StreamContent{stream, content}

}
