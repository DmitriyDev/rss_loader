package rss_loader

import (
	"path/filepath"
	"os"
)

type Loader struct {
	configFile string
}

func (l *Loader) setupConfig() {
	globalConfig = getConfig(l.configFile)
}

func (l *Loader) process() {
	reader := RssReaderStrategy{
		communicator: DZoneCommunicator{globalConfig},
	}

	streamContents := reader.contentOfStreams()

	storage := FileStorage{globalConfig}

	for _, streamContent := range streamContents {
		storage.save(streamContent.stream.Code, streamContent.content)
	}

}

func getWorkDir(absPath bool) string {

	if absPath {
		return filepath.Dir(os.Args[0])
	}

	return filepath.Dir(".")

}
