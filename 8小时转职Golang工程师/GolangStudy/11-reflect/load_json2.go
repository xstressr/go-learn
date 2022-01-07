package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

type MainConfig struct {
	Port string `json:"port"`
	Address string `json:"address"`
}

//反射获取字段中的tag
func reflectTag(mg MainConfig)  {
	mgType := reflect.TypeOf(mg)
	tag0 := mgType.Field(0).Tag
	tag1 := mgType.Field(0).Tag
	fmt.Println(tag0,tag1)
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
	bytes, err := json.Marshal(mainConfig)
	fmt.Println(string(bytes))
	return mainConfig
}

func main() {
	LoadConfig("./test.json")
	config := MainConfig{"1234", "123.221.134"}
	reflectTag(config)

}