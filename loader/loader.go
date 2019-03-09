package loader

type LoaderProcess interface {
	setupConfig()
	process()
}

type Loader struct {
	ConfigFile string
}

func (l Loader) SetupConfig() {
	globalConfig = getConfig(l.ConfigFile)
}

func (l Loader) Process() {
	reader := RssReaderStrategy{
		communicator: DZoneCommunicator{globalConfig},
	}

	streamContents := reader.contentOfStreams()

	storage := FileStorage{globalConfig}

	for _, streamContent := range streamContents {
		storage.save(streamContent.stream.Code, streamContent.content)
	}

}
