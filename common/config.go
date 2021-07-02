package common

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Service struct {
		Name string `yaml:"name"`
		Desc string `yaml:"desc"`
	} `yaml:"service"`
}

var (
	config_     Config
	initialized = false
)

func GetCfg() Config {
	if initialized {
		return config_
	}
	configDir := fmt.Sprintf("%s%s%c%s", GetWorkDir(), "config", os.PathSeparator, "config.yml")
	f, err := os.Open(configDir)
	if err != nil {
		log.Fatal("打开配置文件失败", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config_)
	if err != nil {
		log.Fatal("解析配置文件失败", err)
	}
	initialized = true
	return config_
}
