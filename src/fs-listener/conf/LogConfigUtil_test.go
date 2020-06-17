package conf

import (
	"log"
	"testing"
)

func TestGetLogConf(t *testing.T) {
	t.Run("获取watchDog配置信息", func(t *testing.T) {
		conf := GetLogConf()
		log.Printf("conf 的内存地址：%p \n", &conf)
		emails := conf.Emails
		for _, email := range emails {
			log.Printf("邮箱信息：%v \n", email)
		}
		paths := conf.LogPaths
		for _, logPath := range paths {
			log.Printf("日志地址信息：%v \n", logPath)
		}

		userIds := conf.UserIds
		for _, userId := range userIds {
			log.Printf("用户地址信息：%v \n", userId)
		}

		log.Printf("enable %v \n", conf.Enable)

	})
}
