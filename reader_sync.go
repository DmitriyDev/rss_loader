package main

type SyncRssReader struct {
	config       Config
	communicator Communicator
}

func (srr *SyncRssReader) contentOfStreams(streams map[string]RssStream) []StreamContent {

	var streamsContent []StreamContent

	for _, stream := range streams {
		content := srr.syncRead(stream)
		streamsContent = append(streamsContent, StreamContent{stream, content})
	}

	return streamsContent
}

func (srr *SyncRssReader) syncRead(stream RssStream) string {

	streamUrl := stream.getUrl(srr.config)
	content := srr.communicator.content(streamUrl)
	return content
}
