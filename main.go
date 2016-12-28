package main

import (
	"flag"
	"log"
	"os"

	"github.com/chenxiaoli/crawler/storage"
)

var (
	configFile = flag.String("configfile", "config.ini", "General configuration file")
)

//topic list
var TOPIC = make(map[string]string)

func main() {
	InitCnf()
	storage.StartUp()
	command := os.Args[1]

	log.Println(os.Args[1])

	switch command {

	case "start-worker":
		StartPageCrawlWorker()

	default:
		log.Println("not command input")
	}

}
