package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

type Communicator interface {
	content(streamUrl string) string
	responseBody(resp *http.Response) string
}

type DZoneCommunicator struct {
	config Config
}

func (c DZoneCommunicator) content(streamUrl string) string {

	response, err := http.Get(streamUrl)

	if err != nil {
		log.Fatal(err)
	}

	return c.responseBody(response)
}

func (c DZoneCommunicator) responseBody(resp *http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
