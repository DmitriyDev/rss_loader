package main

type RssReader interface {
	read(map[string]RssStream)
}
