package mailutil

import (
	"cn.com.bluemoon.server.monitor/conf"
	"fmt"
	"gopkg.in/gomail.v2"
)

func GetEmailList() []string {
	alarmConf := conf.GetAlarmConf()
	if nil == alarmConf {
		fmt.Println("获取发送邮件配置失败")
		//
		return []string{"niejian@bluemoon.com.cn"}
	}
	emails := alarmConf.Emails
	return emails
}

func SendMail(message string)  {
	m := gomail.NewMessage()
	m.SetHeader("From", "niejian9001@163.com")
	list := GetEmailList()
	if nil == list || len(list) == 0 {
		// 默认账号
		list = []string{"niejian@bluemoon.com.cn"}
	}

	// slice转可变参数
	m.SetHeader("To", list...)

	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "系统资源告警")
	m.SetBody("text/html", message )
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.163.com", 25, "niejian9001@163.com", "ZQTHKDYFAZOUITOG")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	
}
