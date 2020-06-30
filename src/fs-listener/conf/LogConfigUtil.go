package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// 获取配置信息
func GetLogConf()  *LogConf {
	conf := &LogConf{}
	file, err := ioutil.ReadFile("watchDog.yaml")
	if err != nil {
		return &LogConf{
			LogPaths: []string{},
			Emails:   []string{},
			UserIds:  []string{},
			Errs:     []string{},
			Enable:   false,
			EnableLogPattern: false,
			LogDatePattern: "YYYY-MM-DD",
		}
		// 未找到文件
		log.Println("读取yaml失败", err)
	}

	err = yaml.Unmarshal(file, conf)
	if err != nil {
		log.Fatalf("转换失败 %v \n", err)
	}


	return conf
}
