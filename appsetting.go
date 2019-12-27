package glog

import (
	"log"

	"github.com/BurntSushi/toml"
)

var Config AppSetting

func init() {
	_, err := toml.DecodeFile("appsetting.toml", &Config)
	if err != nil {
		log.Println(err)
	}
}

type AppSetting struct {
	LogService struct {
		Appid   string
		Address string
	}
}