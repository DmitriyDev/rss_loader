package main

type RssReader interface {
	contentOfStreams(streams map[string]RssStream) []StreamContent
}

type StreamContent struct {
	stream RssStream
	content string
}