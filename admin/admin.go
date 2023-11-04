package admin

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AppToken string `yaml:"appToken"`
	Users    []User `yaml:"users"`
}

type User struct {
	UId    string `yaml:"uId"`
	Cookie string `yaml:"cookie"`
}

func GetConfig() Config {
	conf := Config{}
	data, err := os.ReadFile("checkin.yml")
	if err != nil {
		log.Panicf("读取配置失败")
	}
	yaml.Unmarshal(data, &conf)

	return conf
}

func SetConfig(c Config) error {
	f, err := os.OpenFile("checkin.yml", os.O_WRONLY, 0755)
	if err != nil {
		log.Panicf("打开配置失败")
	}
	defer f.Close()
	confYaml, err := yaml.Marshal(c)
	if err != nil {
		log.Panicf("写入配置失败")
	}
	f.Write(confYaml)
	return err
}
