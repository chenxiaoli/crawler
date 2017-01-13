package config

import (
	"flag"
	"log"
	"runtime"

	"github.com/chenxiaoli/crawler/storage"
	"github.com/larspensjo/config"
)

//topic list
var TOPIC = make(map[string]string)

/*
MongoDB 数据库的配置
*/
var MongoDB = make(map[string]string)

/*
RabbitMQ 的配置
*/
var RabbitMQ = make(map[string]string)

/*
Init 初始化
*/
func InitCnf(configFile *string) {

	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	//set config file std
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	//set config file std End

	//Initialized topic from the configuration
	if cfg.HasSection("Mongodb") {
		section, err := cfg.SectionOptions("Mongodb")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("Mongodb", v)
				if err == nil {
					MongoDB[v] = options
				}
			}
		}
	}
	if cfg.HasSection("RabbitMQ") {
		section, err := cfg.SectionOptions("RabbitMQ")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("RabbitMQ", v)
				if err == nil {
					RabbitMQ[v] = options
				}
			}
		}
	}
	//Initialized topic from the configuration END

	storage.InitMongodb(MongoDB)
}
