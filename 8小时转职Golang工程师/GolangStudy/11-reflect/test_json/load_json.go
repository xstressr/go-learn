package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type MainConfig struct {
	port string
	address string
}

func main()  {
	LoadConfig("test.json")
}

func LoadConfig(path string) *MainConfig {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	mainConfig := &MainConfig{}
	err = json.Unmarshal(buf, mainConfig)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}
	fmt.Println(mainConfig.address,mainConfig.port)

	return mainConfig
}