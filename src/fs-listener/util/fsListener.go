package util
//
//import (
//	"bufio"
//	"bytes"
//	"fmt"
//	"fs-listener/conf"
//	"github.com/fsnotify/fsnotify"
//	"log"
//	"os/exec"
//	"regexp"
//	"strings"
//	"sync"
//	"time"
//)
//
//const (
//	WECHAT_ALARM_URL = "http://wechat.bluemoon.com.cn/wxplatform/qyweixin/msg-push/push-msg"
//)
//
//var wg sync.WaitGroup
//
//
//
//func convertWechatMsg()  *conf.Msg{
//	msgText := &conf.MsgText{
//		Content: "",
//	}
//
//	msgData := &conf.MsgData{
//		Touser:  "",
//		MsgType: "text",
//		Agentid: 1000079,
//		Text:    msgText,
//	}
//
//	msg := &conf.Msg{
//		CorpId:  "wx36ef368cf28caea0",
//		Agentid: 1000079,
//		Data:    msgData,
//	}
//
//	return msg
//
//}
//
//// 发送告警
//func sendWxchatAlarm(userIds, alarmMsg string)  {
//	msgText := &conf.MsgText{
//		Content: alarmMsg,
//	}
//
//	msgData := &conf.MsgData{
//		Touser:  userIds,
//		MsgType: "text",
//		Agentid: 1000079,
//		Text:    msgText,
//	}
//
//	msg := &conf.Msg{
//		CorpId:  "wx36ef368cf28caea0",
//		Agentid: 1000079,
//		Data:    msgData,
//	}
//
//	// 发送告警信息
//	resp := Post(WECHAT_ALARM_URL, &msg, "")
//	log.Println("发送微信消息结果：", resp)
//}
//
//
//
//// 监听文件变化
//func GetFsChange(filePath string, errs, emails, userIds []string)  {
//	watcher, err := fsnotify.NewWatcher()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer watcher.Close()
//	// 异常处理
//	defer func() {
//		err := recover()
//		if nil != err {
//			log.Fatalf(" %v \n", err)
//		}
//
//	}()
//
//	done := make(chan bool)
//
//	go func() {
//		for {
//			select {
//			case event := <-watcher.Events:
//				log.Printf("event: %v\n", event)
//
//				// 文件修改事件
//				if event.Op & fsnotify.Write == fsnotify.Write{
//					log.Printf("modifed file %v \n", event)
//					//readChangeContent(event.Name, errs, emails, userIds)
//					fileName := event.Name
//					content := readChangeContent(fileName, errs, emails, userIds)
//					fmt.Printf("读取变化的行数据：%v \n", content)
//					//judgeError(fileName, content, errs, emails, userIds)
//				}
//			case err := <-watcher.Errors:
//				log.Printf("err: %v", err)
//			}
//		}
//	}()
//
//	err = watcher.Add(filePath)
//
//	if err != nil {
//		log.Fatal(err)
//		panic(err)
//	}
//	<-done
//
//}
//
//// 判断是否有异常并且发送警告
//func judgeError(filePath, line string, errs ,emails, userIds []string)  {
//	// 输出异常栈(异常起始为 时间戳 2020-06-01)
//	match := isDatePrefix(line)
//	for _, errTag := range errs {
//		if strings.Contains(line, errTag) && match {
//			msg := line
//			for {
//				match := isDatePrefix(line)
//				msg = convertAlarmMsg(msg, line)
//
//				if match {
//					break
//				}
//
//			}
//			userIds := strings.Join(userIds, "|")
//			// 发送告警信息
//			go sendWxchatAlarm(userIds, msg)
//		}
//	}
//}
//
//// 判断字符串是否是日期时间戳开头
//func isDatePrefix(line string) bool {
//	r := []rune(line)
//	newLine20Prefix := string(r[0:19])
//	pattern :="\\d{4}\\-\\d{2}\\-\\d{2}\\s\\d{2}:\\d{2}:\\d{2}"
//	match, _ := regexp.Match(pattern, []byte(newLine20Prefix))
//	return match
//}
//
//func convertAlarmMsg(msg, line string) string {
//	return msg + "\n" + line
//}
//
//func analyseAndAlarm(line string) string {
//	msg := ""
//	match := isDatePrefix(line)
//
//	// 错误堆栈信息不是时间戳开头
//	if !match {
//		msg = convertAlarmMsg(msg, line)
//	}
//	return msg
//}
//
//
//
//var alarmContentChan = make(chan string)
//
//func getalarmContent()  {
//	alarmContent := <-alarmContentChan
//	fmt.Printf("堆栈信息：%v\n", alarmContent)
//}
//
//// 读取文件变化的内容
//func readChangeContent(file string, errs ,emails, userIds []string) string{
//	// chan 信道
//	var errContentChan = make(chan string)
//	fmt.Println("监控文件：", file)
//	newLine := ""
//	// 每次指定文件的读取一行数据
//	command := exec.Command("sh", "-c", "tail -1f " + file)
//	stdoutPipe, err := command.StdoutPipe()
//	if err != nil {
//		panic(err)
//	}
//
//	//quit := make(chan bool)
//
//	// 读取信道数据
//	go func() {
//		println("========channel 写入数据======")
//		//var msg = ""
//
//		for  {
//			select {
//			case <- time.After(2 * time.Second):
//				fmt.Println("timeout write")
//				// 2s未写，发送退出信号
//				//quit<-true
//				var buffer bytes.Buffer
//				for con := range errContentChan {
//					//fmt.Printf("content：%v\n", con)
//					buffer.WriteString(con)
//					//buffer.WriteString("\n")
//					//msg = buffer.String()
//					//fmt.Printf("堆栈信息1：%v\n", buffer.String())
//				}
//				fmt.Printf("堆栈信息1：%v\n", buffer.String())
//				alarmContentChan <- buffer.String()
//			//case con := <-errContentChan:
//			//	fmt.Println("收到堆栈信息：", con)
//			case errContentChan <- "":
//				//fmt.Println("写入errContentChan")
//
//
//			}
//			if len(alarmContentChan) > 0 {
//				fmt.Println("alarmContentChan:", len(alarmContentChan))
//				break
//			}
//		}
//
//		getalarmContent()
//
//
//		//go sendWxchatAlarm(userIds, msg)
//	}()
//
//	go func() {
//		reader := bufio.NewReader(stdoutPipe)
//
//		for {
//			// 堆栈已经抛完，关闭信道，通知信道读取数据
//			//if len(errContentChan) >=10 {
//			//	close(errContentChan)
//			//	continue
//			//}
//			fmt.Println("-----------")
//			line, _, err2 := reader.ReadLine()
//			//newLine, err2 := reader.ReadString('\n')
//
//			//fmt.Printf("读取变化的行数据：prefix：%v", prefix)
//			if err2 != nil {
//				panic(err2)
//			}
//
//			//fmt.Println("读到数据：", newLine)
//
//			newLine = string(line)
//			fmt.Println("读到数据：", newLine)
//			if strings.TrimSpace(newLine) == " " {
//				continue
//			}
//			match := isDatePrefix(newLine)
//			hasExp := false
//			//fmt.Printf("newLine: %v \n", newLine)
//
//			for _, errTag := range errs {
//				// 含有异常关键字
//				if strings.Contains(newLine, errTag) {
//					hasExp = true
//					break
//				}
//			}
//
//			if hasExp && match {
//				errContentChan <- newLine
//			}
//
//			if hasExp && !match {
//				errContentChan <- newLine
//			}
//
//			if !hasExp && !match {
//				errContentChan <- newLine
//			}
//
//			//select {
//			//case <-quit:
//			//	fmt.Println("写超时2S，关闭errContentChan")
//			//	close(errContentChan)
//			//default:
//			//	fmt.Println("quit 默认操作")
//			//
//			//}
//			//fmt.Printf("chan 长度：%v\n", len(errContentChan))
//
//		}
//	}()
//
//	err = command.Run()
//	if err != nil {
//		panic(err)
//	}
//
//	return newLine
//}
//
////func readChangeContent(file string, errs ,emails, userIds []string)  {
////	fmt.Println("监控文件：", file)
////	// 每次指定文件的读取一行数据
////	command := exec.Command("sh", "-c", "tail -1f "+file)
////	stdoutPipe, err := command.StdoutPipe()
////	if err != nil {
////		panic(err)
////	}
////	go func() {
////		reader := bufio.NewReader(stdoutPipe)
////		for {
////			line, _, err2 := reader.ReadLine()
////			//fmt.Printf("读取变化的行数据：prefix：%v", prefix)
////			if err2 != nil {
////				panic(err2)
////			}
////			newLine := string(line)
////			// 输出异常栈(异常起始为 时间戳 2020-06-01)
////			r := []rune(newLine)
////			newLine20Prefix := string(r[0:19])
////
////			// 判断输出的行前20位是否是时间戳的格式，如果是 2020-06-16 09:59:24
////			//pattern :="\\d{4}-\\d{4}-\\d{4}-\\ \\s \\d{2}:\\d{2}:\\d{2}:"
////			pattern :="\\d{4}\\-\\d{2}\\-\\d{2}\\s\\d{2}:\\d{2}:\\d{2}"
////			match, _ := regexp.Match(pattern, []byte(newLine20Prefix))
////			//fmt.Printf("字符串：%v, 是否是时间戳格式：%v \n", newLine20Prefix, match)
////
////			for {
////				for _, errTag := range errs {
////					// 包含敏感的错误信息，发送告警通知
////					if strings.Contains(newLine, errTag) {
////
////
////						userIds := strings.Join(userIds, "|")
////						// 发送告警信息
////						go sendWxchatAlarm(userIds, newLine)
////					}
////				}
////			}
////
////		}
////	}()
////	err = command.Run()
////	if err != nil {
////		panic(err)
////	}
////
////}