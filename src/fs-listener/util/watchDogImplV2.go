package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
	"watchdog/conf"

	"github.com/fsnotify/fsnotify"
	"github.com/hpcloud/tail"
	"github.com/patrickmn/go-cache"
)

const (
	WECHAT_ALARM_URL = "http://wechat.bluemoon.com.cn/wxplatform/qyweixin/msg-push/push-msg"
	DATE_TAG         = "$DATE$"
	ERROR_TAG        = "ERROR"
	DEBUG_TAG = "DEBUG"
	WARN_TAG = "WARN"
)

var lock sync.Mutex

// 发送告警
func sendWxchatAlarm(userIds, alarmMsg string) {
	msgText := &conf.MsgText{
		Content: alarmMsg,
	}

	msgData := &conf.MsgData{
		Touser:  userIds,
		MsgType: "text",
		Agentid: 1000079,
		Text:    msgText,
	}

	msg := &conf.Msg{
		CorpId:  "wx36ef368cf28caea0",
		Agentid: 1000079,
		Data:    msgData,
	}

	// 发送告警信息
	go Post(WECHAT_ALARM_URL, &msg, "")
}

func getFileList(path string) []string {
	logPathList := make([]string, 10)
	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		if file.IsDir() {
			log.Println(path + file.Name())
			getFileList(path + file.Name() + "/")
		} else {
			logName := path + file.Name()
			//pattern :="\\d{4}\\-\\d{2}\\-\\d{2}\\s\\d{2}:\\d{2}:\\d{2}"
			pattern := "\\d{4}\\-\\d{2}\\-\\d{2}\\-\\d"
			match, _ := regexp.Match(pattern, []byte(logName))

			if !match {
				if strings.Contains(logName, ".log") {
					logPathList = append(logPathList, logName)
				}

			}

		}
	}

	return logPathList
}

var fileMap sync.Map

// 监听文件变化
func GetFsChange(filePath string, errs, emails, userIds []string, cache *cache.Cache) {
	done := make(chan bool)
	log.Printf("日志文件：%v \n", filePath)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("err: ", err)
	}

	// 异常处理
	defer func() {
		err := recover()
		if nil != err {
			log.Fatalf("捕获异常: %v \n", err)
		}

	}()
	defer watcher.Close()
	//quit := make(chan bool)

	err = watcher.Add(filePath)

	if err != nil {
		log.Fatal("err===>:", err)
		panic(err)
	}
	

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				//log.Printf("event: %v\n", event)
				fileName := event.Name

				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("创建文件 : ", fileName)
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					//log.Println(strings.HasSuffix(fileName, ".log"))
					// 只读 .log, .out的日志信息
					if !strings.HasSuffix(fileName, ".log") && !strings.HasSuffix(fileName, ".out") {
						log.Printf("文件 %v， 可能是归档文件，无需处理 \n", fileName)
						continue
					}

					log.Println("写入文件 : ", fileName)
					_, ok := fileMap.Load(fileName)
					fmt.Printf("%v 是否被监听：%v \n", fileName, ok)

					if !ok {
						fileMap.Store(fileName, fileName)
						go readChangeContent(fileName, errs, emails, userIds, cache)
					} else {
						fmt.Printf("文件已被监听：%v ，无需再注册监听 \n", fileName)
					}

				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("删除文件 : ", fileName)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Println("重命名文件 : ", fileName)
				}
				if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					log.Println("修改权限 : ", fileName)
				}

			case err := <-watcher.Errors:
				log.Printf("err: %v", err)
			}
		}
	}()
	<-done
}

// 判断字符串是否是日期时间戳开头
func isDatePrefix(line string) bool {
	r := []rune(line)
	newLine20Prefix := string(r[0:19])
	pattern := "\\d{4}\\-\\d{2}\\-\\d{2}\\s\\d{2}:\\d{2}:\\d{2}"
	match, _ := regexp.Match(pattern, []byte(newLine20Prefix))
	return match
}

// 组合告警信息
// 告警应用
// 工具IP
func convertAlarmMsg(err, line, logFilePath string) string {
	//logFilePaths := strings.Split(logFilePath, "/")
	//appName := ""
	//length := len(logFilePaths)
	//if len(logFilePaths) > 0  {
	//	appName = logFilePaths[length-1]
	//}
	ip := GetNetIp()
	var buffer bytes.Buffer
	buffer.WriteString("IP：" + ip + "<br/>")
	buffer.WriteString("时间：" + FormatDate("2006-01-02 15:04:05") + "<br/>")
	buffer.WriteString("日志：" + logFilePath + "<br/>")
	buffer.WriteString("敏感异常：" + err + "<br/>")
	buffer.WriteString("堆栈信息<br/>")
	buffer.WriteString(line)

	return buffer.String()
}

var alarmContentChan = make(chan string)

func FormatDate(pattern string) string {
	return time.Now().Format(pattern)
}

func IsToday(day string) bool {
	date := FormatDate("2006-01-02")
	return day == date
}

func getalarmContent() {
	alarmContent := <-alarmContentChan
	fmt.Printf("堆栈信息：%v\n", alarmContent)
}

func cacheAndSendMsg(msg, file, custErrs string, myCache *cache.Cache, userIds []string, )  {

	if "" == strings.TrimSpace(msg) {
		fmt.Printf("信息内容为空\n")
		return
	}

	md5Str := md52Str(msg)
	_, isExist := myCache.Get(md5Str)
	if !isExist {
		// 防止重复消息发送
		myCache.Set(md5Str, md5Str, cache.DefaultExpiration)
		users := strings.Join(userIds, "|")
		log.Println("发送消息:", users)
		sendWxchatAlarm(users, convertAlarmMsg(custErrs, msg, file))
	} else {
		log.Printf("已发送消息 %v \n", md5Str)
	}
}

// 读取文件变化的内容
func readChangeContent(file string, errs, emails, userIds []string, myCache *cache.Cache) string {
	newLine := ""
	// 异常处理
	defer func() {
		if r := recover();r!=nil{
			log.Printf("tail err: %v",r)
		}
	}()


	tailfs, err := tail.TailFile(file, tail.Config{
		ReOpen:    true,                                 // 文件被移除或被打包，需要重新打开
		Follow:    true,                                 // 实时跟踪
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 如果程序出现异常，保存上次读取的位置，避免重新读取。 whence，从哪开始：0从头，1当前，2末尾
		MustExist: false,                                // 如果文件不存在，是否推出程序，false是不退出
		Poll:      true,
	})

	if err != nil {
		log.Printf("tailf failed, err:%v \n", err)
		panic(err)
	}

	//errContentChan := make(chan string)
	//writing := make(chan bool)
	quit := make(chan bool)
	msg := ""

	for line := range tailfs.Lines {
		//fmt.Println("读到信息->：", line.Text)
		newLine = line.Text
		//loc, _ := time.LoadLocation("Asia/Shanghai")

		//tailTime := line.Time.In(loc)
		//fmt.Printf("tailTime---> %v \n", tailTime)

		match := isDatePrefix(newLine)
		hasExp := false
		custErr := ""



		for _, errTag := range errs {
			//fmt.Printf("errTag：%v, newLine: %v \n", errTag, newLine)
			// 含有异常关键字，发送提示告警
			if strings.Contains(newLine, errTag) {
				custErr = errTag
				hasExp = true
				break
			}

		}

		// log.error 输出方式，
		if !hasExp && match && strings.Contains(newLine, ERROR_TAG) {
			//fmt.Println("msg：", msg)
			if "" != msg  {
				////msgHasExp := false
				//for _, errTag := range errs {
				//	//fmt.Printf("errTag：%v, newLine: %v \n", errTag, newLine)
				//	// 含有异常关键字，发送提示告警
				//	if strings.Contains(msg, errTag) {
				//		custErr = errTag
				//		hasExp = true
				//		break
				//	}
				//
				//}
				//if hasExp {
				//	msg += newLine + "\n"
				//	//cacheAndSendMsg(msg, file, custErr, myCache, userIds)
				//}
			} else {
				// msg 为空
				if !strings.Contains(newLine, DEBUG_TAG) || !strings.Contains(newLine, WARN_TAG) {
					msg += newLine + "\n"
				}
			}
		}

		if hasExp && match {
			//var sendMsgFlag = false
			//fmt.Println("msg：", msg)
			// 读到下一个错误，或是错误开头
			//if "" != msg {
			//	sendMsgFlag = true
			//
			//}

			//errContentChan <- newLine + "\n"
			//writing <- true
			if !strings.Contains(newLine, DEBUG_TAG) || !strings.Contains(newLine, WARN_TAG) {
				msg += newLine + "\n"
			}

			//if sendMsgFlag {
			//	cacheAndSendMsg(msg, file, custErr, myCache, userIds)
			//}
			// msg 置空
			//msg = ""
		}

		if hasExp && !match {
			//errContentChan <- newLine + "\n"
			//writing <- true
			msg += newLine + "\n"
		}

		if !hasExp && !match {
			//errContentChan <- newLine + "\n"
			//writing <- true
			msg += newLine + "\n"
		}

		// 0.5秒没操作，判断是需要发送消息
		time.AfterFunc(500 * time.Millisecond, func() {
			//fmt.Println("时间静止500MS")
			if "" != msg{
				for _, errTag := range errs {
					//fmt.Printf("errTag：%v, newLine: %v \n", errTag, newLine)
					// 含有异常关键字，发送提示告警
					if strings.Contains(msg, errTag) {
						custErr = errTag
						hasExp = true
						break
					}

				}
				if hasExp {
					lock.Lock()
					cacheAndSendMsg(msg, file, custErr, myCache, userIds)
					lock.Unlock()
				}
			}
			msg = ""
		})


	}

	<-quit

	return newLine
}

func md52Str(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
