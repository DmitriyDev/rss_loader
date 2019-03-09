package loader

import "fmt"

type RssReaderStrategy struct {
	communicator Communicator
}

type Scenario interface {
	contentOfStreams(streams map[string]RssStream) []StreamContent
}

type StreamContent struct {
	stream  RssStream
	content string
}

func (r *RssReaderStrategy) contentOfStreams() []StreamContent {

	var readScenario Scenario

	if globalConfig.Async {
		readScenario = r.asyncScenario()
	} else {
		readScenario = r.syncScenario()
	}

	return readScenario.contentOfStreams(globalConfig.Streams)
}

func (r *RssReaderStrategy) syncScenario() *SyncScenario {

	fmt.Println("Sync mode")

	return &SyncScenario{
		communicator: r.communicator,
	}

}

func (r *RssReaderStrategy) asyncScenario() *AsyncScenario {

	fmt.Println("Async mode")

	return &AsyncScenario{
		communicator: r.communicator,
	}

}
