package main

import (
	"fmt"
	"time"
	lr "github.com/DmitriyDev/rss_loader/loader"
)

const CONFIG_FILE = "/config.yml"

func init() {
	fmt.Println("Start process")
}

func main() {
	start := time.Now()

	l := lr.Loader{ConfigFile: CONFIG_FILE, }
	l.SetupConfig()
	l.Process()

	fmt.Println(time.Since(start).Seconds())

}

