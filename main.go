package main

import (
	"fmt"

	"github.com/JerryCheese/dlems/store"
)

var appConf AppConf

func main() {

	conf, err := loadConfig("env.yaml")
	if err != nil {
		fmt.Printf("fail to load config. %v\n", err)
		return
	}
	appConf = conf

	store.Init(conf.Mongo)

	initAPI(conf.API)
}
