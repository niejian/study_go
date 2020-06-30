package util
//
//import (
//	"bufio"
//	"bytes"
//	"crypto/md5"
//	"encoding/hex"
//	"fmt"
//	"github.com/fsnotify/fsnotify"
//	"github.com/patrickmn/go-cache"
//	"io/ioutil"
//	"log"
//	"os/exec"
//	"regexp"
//	"strings"
//	"sync"
//	"time"
//	"watchdog/conf"
//)
//
//const (
//	WECHAT_ALARM_URL = "http://wechat.bluemoon.com.cn/wxplatform/qyweixin/msg-push/push-msg"
//	DATE_TAG = "$DATE$"
//	ERROR_TAG = "ERROR"
//)
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
//func getFileList(path string) []string{
//	logPathList := make([]string, 10)
//	fs,_:= ioutil.ReadDir(path)
//	for _,file:=range fs{
//		if file.IsDir(){
//			log.Println(path+file.Name())
//			getFileList(path+file.Name()+"/")
//		}else{
//			logName := path+file.Name()
//			//pattern :="\\d{4}\\-\\d{2}\\-\\d{2}\\s\\d{2}:\\d{2}:\\d{2}"
//			pattern :="\\d{4}\\-\\d{2}\\-\\d{2}\\-\\d"
//			match, _ := regexp.Match(pattern, []byte(logName))
//
//			if !match {
//				if strings.Contains(logName, ".log"){
//					logPathList = append(logPathList, logName)
//				}
//
//			}
//
//		}
//	}
//
//	return logPathList
//}
//
//var fileMap = make(map[string]string)
//
//// 监听文件变化
//func GetFsChange(filePath string, errs, emails, userIds []string, cache *cache.Cache)  {
//	done := make(chan bool)
//	log.Printf("日志文件：%v \n", filePath)
//	watcher, err := fsnotify.NewWatcher()
//	if err != nil {
//		log.Fatal("err: ", err)
//	}
//
//
//	// 异常处理
//	defer func() {
//		err := recover()
//		if nil != err {
//			log.Fatalf("捕获异常: %v \n", err)
//		}
//
//	}()
//	defer watcher.Close()
//	//quit := make(chan bool)
//
//	err = watcher.Add(filePath)
//
//	if err != nil {
//		log.Fatal("err===>:", err)
//		panic(err)
//	}
//
//	var fileMap sync.Map
//
//	go func() {
//		for {
//			select {
//			case event := <-watcher.Events:
//				//log.Printf("event: %v\n", event)
//				fileName := event.Name
//
//				if event.Op&fsnotify.Create == fsnotify.Create {
//					log.Println("创建文件 : ", fileName);
//				}
//				if event.Op&fsnotify.Write == fsnotify.Write {
//
//					log.Println("写入文件 : ", fileName);
//					_, ok := fileMap.Load(fileName)
//					if !ok {
//						go readChangeContent(fileName, errs, emails, userIds, cache)
//					} else {
//						fileMap.Store(fileName, fileName)
//					}
//
//
//
//				}
//				if event.Op&fsnotify.Remove == fsnotify.Remove {
//					log.Println("删除文件 : ", fileName);
//				}
//				if event.Op&fsnotify.Rename == fsnotify.Rename {
//					log.Println("重命名文件 : ", fileName);
//				}
//				if event.Op&fsnotify.Chmod == fsnotify.Chmod {
//					log.Println("修改权限 : ", fileName);
//				}
//
//			case err := <-watcher.Errors:
//				log.Printf("err: %v", err)
//			}
//		}
//	}()
//	<-done
//	select {
//
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
//// 组合告警信息
//// 告警应用
//// 工具IP
//func convertAlarmMsg(err, line, logFilePath string) string {
//	//logFilePaths := strings.Split(logFilePath, "/")
//	//appName := ""
//	//length := len(logFilePaths)
//	//if len(logFilePaths) > 0  {
//	//	appName = logFilePaths[length-1]
//	//}
//	ip := GetNetIp()
//	var buffer bytes.Buffer
//	buffer.WriteString("IP：" + ip+ "<br/>")
//	buffer.WriteString("时间：" + FormatDate("2006-01-02 15:04:05") + "<br/>")
//	buffer.WriteString("日志：" + logFilePath + "<br/>")
//	buffer.WriteString("敏感异常：" + err + "<br/>")
//	buffer.WriteString("堆栈信息<br/>" )
//	buffer.WriteString(line)
//
//	return buffer.String()
//}
//
//var alarmContentChan = make(chan string)
//
//func FormatDate(pattern string) string  {
//	return time.Now().Format(pattern)
//}
//
//func getalarmContent()  {
//	alarmContent := <-alarmContentChan
//	fmt.Printf("堆栈信息：%v\n", alarmContent)
//}
//
//// 读取文件变化的内容
//func readChangeContent(file string, errs ,emails, userIds []string, myCache *cache.Cache) string{
//	newLine := ""
//	// 每次指定文件的读取一行数据
//	command := exec.Command("sh", "-c", "tail -1f " + file)
//	stdoutPipe, err := command.StdoutPipe()
//	if err != nil {
//		panic(err)
//	}
//
//	errContentChan := make(chan string)
//	//writing := make(chan bool)
//	quit := make(chan bool)
//
//	//sendMsgFlag := make(chan bool)
//	//doMsg := make(chan bool)
//	//custErr := ""
//	go func() {
//		reader := bufio.NewReader(stdoutPipe)
//		for {
//
//			line, _, err2 := reader.ReadLine()
//
//			if err2 != nil {
//				panic(err2)
//			}
//
//			newLine = string(line)
//			//fmt.Println("读到数据：", newLine)
//			if strings.TrimSpace(newLine) == " " {
//				continue
//			}
//
//			//fmt.Printf("newLine: %v \n", newLine)
//			match := isDatePrefix(newLine)
//			hasExp := false
//
//			for _, errTag := range errs {
//				// 含有异常关键字，发送提示告警
//				if strings.Contains(newLine, errTag) {
//					//custErr = errTag
//					hasExp = true
//					break
//				}
//
//			}
//
//			if !hasExp && match && strings.Contains(newLine, ERROR_TAG) {
//				errContentChan <- newLine + "\n"
//			}
//
//			if hasExp && match {
//				errContentChan <- newLine + "\n"
//				//writing <- true
//			}
//
//			if hasExp && !match {
//				errContentChan <- newLine + "\n"
//				//writing <- true
//			}
//
//			if !hasExp && !match {
//				errContentChan <- newLine + "\n"
//				//writing <- true
//			}
//			command.Run()
//
//			//log.Println("lines--: ", newLine)
//
//			//for con := range errContentChan {
//			//	log.Println(con)
//			//}
//
//		}
//
//	}()
//
//	msg := ""
//	go func() {
//
//		for{
//
//			select {
//				// channel超时, 0.15秒未接收到写要求，直接发送错误信息
//				case <-time.After(150 * time.Millisecond):
//					//log.Println("写入超时 ：", len(errContentChan))
//					custErrs := ""
//					hasCustErr := false
//					if "" != msg {
//						//log.Println("异常堆栈：", msg)
//
//						// msg中可能含有多类异常信息，按照相关的规律分别发送这些异常信息
//						for _, errTag := range errs {
//							hasCustErr = true
//							if strings.Contains(msg, errTag) {
//								custErrs = errTag
//							}
//						}
//
//						if hasCustErr && "" != custErrs {
//							log.Printf("hasCustErr %v \n", hasCustErr)
//							//strings.Split(msg, "20")
//							//sendMsgFlag <- true
//							md5Str := md52Str(msg)
//							_, isExist := myCache.Get(md5Str)
//							if !isExist {
//								myCache.Set(md5Str, md5Str, cache.DefaultExpiration)
//								go sendWxchatAlarm(strings.Join(userIds, "|"), convertAlarmMsg(custErrs, msg, file) )
//							} else {
//								log.Printf("已发送消息 %v \n", md5Str)
//							}
//
//							//<-errContentChan
//						}
//					}
//
//					// 置空
//					msg = ""
//					//done <- true
//				case content := <-errContentChan : // 阻塞，输出
//					// 组装异常堆栈信息
//					msg += content
//					//log.Println("异常堆栈 content：", content)
//				//case <- sendMsgFlag:
//				//	msg = ""
//
//			}
//
//
//		}
//	}()
//
//
//	err = command.Run()
//	if err != nil {
//		panic(err)
//	}
//	<-quit
//
//	return newLine
//}
//
//func md52Str(str string) string  {
//	h := md5.New()
//	h.Write([]byte(str))
//	return hex.EncodeToString(h.Sum(nil))
//}
//
