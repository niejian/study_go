package main

import (
	"fs-listener/conf"
	"fs-listener/util"
	"log"
)

func main() {
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

	for _, path := range paths {
		// 获取配置信息
		util.GetFsChange(path, errs, emails, userIds)
	}
}


