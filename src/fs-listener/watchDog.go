package main

import (
	"fmt"
	"fs-listener/conf"
	"fs-listener/util"
	"log"
	"os"
	"time"
)

func initLog() {
	file := "./watchdog-" + time.Now().Format("2006-01-02") + ".log"


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
	done := make(chan bool)
	for _, path := range paths {
		// 获取配置信息
		go util.GetFsChange(path, errs, emails, userIds)
	}
	<-done
}


