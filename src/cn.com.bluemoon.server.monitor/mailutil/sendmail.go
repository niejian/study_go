package mailutil

import "gopkg.in/gomail.v2"

func SendMail(message string)  {
	m := gomail.NewMessage()
	m.SetHeader("From", "niejian9001@163.com")
	m.SetHeader("To", "niejian@bluemoon.com.cn")
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
