package main

import (
	"thrift-test/wechatmsg"
	"time"
)

func main() {

	// 多线程方式发送消息
	done := make(chan bool)
	for {
		go wechatmsg.SendMsg("80468295", "hello world" + time.Now().Format(time.RFC3339))
	}

	<-done

}

