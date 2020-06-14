package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// 获取配置信息
func  GetAlarmConf() *alarmConf {
	// 申明结构体
	conf := &alarmConf{}
	yamlFile, err := ioutil.ReadFile("serverMonitor.yaml")
	if err != nil {
		log.Printf("读取配置文件失败：%v", err)
		// 给出默认值
		return &alarmConf{
			120.0,
			110.0,
			100.0,
			80.0,
			[]string{"niejian@bluemoon.com.cn"},
		}
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return conf
}