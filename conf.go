package main

import (
	"io/ioutil"

	"github.com/JerryCheese/dlems/store"
	"gopkg.in/yaml.v2"
)

// AppConf _
type AppConf struct {
	API   APIConf         `yaml:"api"`
	Mongo store.MongoConf `yaml:"mongo"`
}

// loadConfig _
func loadConfig(confFile string) (conf AppConf, err error) {
	b, err := ioutil.ReadFile(confFile)
	if err != nil {
		return
	}
	yaml.Unmarshal(b, &conf)
	return
}
