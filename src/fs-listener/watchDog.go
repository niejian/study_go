package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"watchdog/conf"
	"watchdog/util"

	"github.com/patrickmn/go-cache"
)

func initLog() {
	file := "./watchdog.log"

	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}

	//创建一个Logger
	//参数1：日志写入目的地
	//参数2：每条日志的前缀
	//参数3：日志属性
	log.New(logFile, "前缀", log.Ldate|log.Ltime|log.Lshortfile)

	//Flags返回Logger的输出选项
	fmt.Println(log.Flags())
	//SetFlags设置输出选项
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)

	//返回输出前缀
	fmt.Println(log.Prefix())

}

func main() {
	initLog()
	logConf := conf.GetLogConf()
	if !logConf.Enable {
		log.Println("已关闭告警")
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
	logDatePattern := logConf.LogDatePattern     // 监控日志文件的日期格式

	done := make(chan bool)

	for _, path := range paths {
		log.Printf("enableLogPattern : %v,  logDatePattern: %v, path: %v\n", enableLogPattern, logDatePattern, path)
		isContinue := false
		// 带有日期的日志文件
		if enableLogPattern && strings.Contains(path, util.DATE_TAG) && "" != logDatePattern {
			var date = ""
			// 读取日期格式配置
			if strings.Contains(logDatePattern, "-") {
				date = util.FormatDate("2006-01-02")
			} else {
				date = util.FormatDate("20060102")
			}

			path = strings.ReplaceAll(path, util.DATE_TAG, date)
			isContinue = true
		}

		if !strings.Contains(path, util.DATE_TAG) {
			isContinue = true
		}

		if isContinue {
			exists, _ := util.PathExists(path)
			if !exists {
				log.Printf("文件 %v \n不存在", path)
				continue
			}
		}

		if isContinue {
			log.Printf("监控文件-->：%v,  \n", path)
			// 过期时间30s， 每隔20秒清除过期的key
			c := cache.New(30*time.Second, 20*time.Second)

			go util.GetFsChange(path, errs, emails, userIds, c)
		}

	}
	
	<-done
}
