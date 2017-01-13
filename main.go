package main

import (
	"flag"
	"os"

	"github.com/chenxiaoli/crawler/config"
	"github.com/chenxiaoli/crawler/storage"
)

func main() {

	configFile := flag.String("configfile", os.Args[1], "General configuration file")
	config.InitCnf(configFile)
	storage.StartUp()

	StartPageCrawlWorker()

}
