package main

type SyncScenario struct {
	communicator Communicator
}

func (srr *SyncScenario) contentOfStreams(streams map[string]RssStream) []StreamContent {

	var streamsContent []StreamContent

	for _, stream := range streams {
		content := srr.syncRead(stream)
		streamsContent = append(streamsContent, StreamContent{stream, content})
	}

	return streamsContent
}

func (srr *SyncScenario) syncRead(stream RssStream) string {

	streamUrl := stream.getUrl()
	content := srr.communicator.content(streamUrl)
	return content
}
