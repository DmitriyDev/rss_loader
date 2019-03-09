package main

import (
	"fmt"
	"time"
	"github.com/DmitriyDev/rss_loader/loader"
)

const CONFIG_FILE = "/config.yml"

func init() {
	fmt.Println("Start process")
}

func main() {
	start := time.Now()

	l := lr.Loader{}
	l.setupConfig(CONFIG_FILE)
	l.process()

	fmt.Println(time.Since(start).Seconds())

}

