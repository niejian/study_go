package util

import (
	"github.com/patrickmn/go-cache"
	"log"
	"strings"
	"testing"
	"time"
	"watchdog/conf"
)


func TestGetFsChange(t *testing.T) {
	t.Run("文件变化监控", func(t *testing.T) {
		//GetFsChange("/Users/a/logs/demo-muti-registry-producer/")
		logConf := conf.GetLogConf()
		if !logConf.Enable {
			log.Printf("已关闭告警")
			return
		}
		paths := logConf.LogPaths
		if nil == paths || len(paths) == 0 {
			panic("请选择监控的日志路径")
		}

		errs := logConf.Errs
		if nil == errs || len(errs) == 0 {
			panic("未发现可监控的异常信息")
		}

		emails := logConf.Emails
		if nil == emails || len(emails) == 0 {
			emails = []string{"niejian@bluemoon.com.cn"}
		}

		userIds := logConf.UserIds
		if nil == userIds || len(userIds) == 0 {
			userIds = []string{"80468295"}
		}

		enableLogPattern := logConf.EnableLogPattern // 监控文件名是否开启正则表达匹配模式
		logDatePattern := logConf.LogDatePattern // 监控日志文件的日期格式

		done := make(chan bool)


		for _, path := range paths {
			//log.Printf("enableLogPattern : %v,  logDatePattern: %v, path: %v\n", enableLogPattern, logDatePattern, path)
			isContinue := false
			// 带有日期的日志文件
			if enableLogPattern && strings.Contains(path, DATE_TAG) && "" != logDatePattern {

				var date = ""
				// 读取日期格式配置
				if strings.Contains(logDatePattern, "-") {
					date = FormatDate("2006-01-02")
				} else {
					date = FormatDate("20060102")
				}
				path = strings.ReplaceAll(path, DATE_TAG, date)
				isContinue = true
			}

			if !strings.Contains(path, DATE_TAG) {
				isContinue = true
			}

			if isContinue {
				exists, _ := PathExists(path)
				if !exists {
					log.Printf("文件 %v \n不存在", path)
					continue
				}
			}

			if isContinue {
				// 过期时间10s， 每隔五秒清除过期的key
				c := cache.New(10*time.Second, 5*time.Second)

				go GetFsChange(path, errs, emails, userIds, c)
			}


		}
		<-done

	})
}

