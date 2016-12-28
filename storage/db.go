package storage

import (
	"log"

	"fmt"

	"gopkg.in/mgo.v2"
)

var (
	mgoSession *mgo.Session
	dataBase   string
	mongodbURL string
)

/*
InitMongodb 初始化mongodb
*/
func InitMongodb(mgoTopic map[string]string) {
	mongodbURL = fmt.Sprintf("%s:%s", mgoTopic["host"], mgoTopic["port"])
	dataBase = mgoTopic["dbname"]
}

/*
StartUp 启动
*/
func StartUp() {
	if mgoSession == nil {
		var err error
		log.Println("mongodbURL:" + mongodbURL)
		mgoSession, err = mgo.Dial(mongodbURL)
		if err != nil {
			panic(err) //直接终止程序运行
		}
	}
}

/*
 GetSession 公共方法，获取session，如果存在则拷贝一份
*/
func GetSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(mongodbURL)
		if err != nil {
			panic(err) //直接终止程序运行
		}
		log.Println("new connection")
	}
	//最大连接池默认为4096
	return mgoSession
}

//公共方法，获取collection对象
func witchCollection(collection string, s func(*mgo.Collection) error) error {
	session := GetSession()
	defer session.Close()
	c := session.DB(dataBase).C(collection)
	return s(c)
}
