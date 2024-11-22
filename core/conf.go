package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"nft/server/config"
	global "nft/server/gloabl"
)

// initConf 读取yaml文件的配置
func InitConf() {
	const ConfigFile = "settings.yaml" // yaml文件路径
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init unmarshal: %v", err)
	}
	log.Println("config yamlFile load init success.")
	global.Config = c
}

// SetYaml 修改配置文件
//func SetYaml() error {
//	byteData, err := yaml.Marshal(global.Config)
//	if err != nil {
//		return err
//	}
//	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
//	if err != nil {
//		return err
//	}
//	global.Log.Info("修改配置文件成功！")
//	return nil
//}
