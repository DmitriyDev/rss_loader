package loader

import (
	"net/http"
	"io/ioutil"
	"log"
)

type Communicator interface {
	content(streamUrl string) string
}

type DZoneCommunicator struct {
	config Config
}

func (dzc DZoneCommunicator) content(streamUrl string) string {

	response, err := http.Get(streamUrl)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
