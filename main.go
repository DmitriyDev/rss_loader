package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Start process")

	c := getConfig()

	reader := Reader{
		config:       c,
		storage:      FileStorage{c},
		communicator: DZoneCommunicator{c},
	}
	reader.read()

	fmt.Println(time.Since(start).Seconds())

}

