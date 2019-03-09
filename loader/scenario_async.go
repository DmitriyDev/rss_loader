package loader

type AsyncScenario struct {
	communicator Communicator
}

func (arr *AsyncScenario) contentOfStreams(streams map[string]RssStream) []StreamContent {

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

func (arr *AsyncScenario) asyncRead(stream RssStream, ch chan StreamContent) {

	streamUrl := stream.getUrl()
	content := arr.communicator.content(streamUrl)
	ch <- StreamContent{stream, content}

}
