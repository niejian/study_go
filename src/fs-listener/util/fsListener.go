package util

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os/exec"
)

// 监听文件变化

func GetFsChange(filePath string)  {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()
	// 异常处理
	defer func() {
		err := recover()
		if nil != err {
			log.Fatalf("读取文件失败")
		}
		
	}()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Printf("event: %v\n", event)
				// 文件修改事件
				if event.Op & fsnotify.Write == fsnotify.Write{
					log.Printf("modifed file %v \n", event)
					go readChaneContent(event.Name)
				}
			case err := <-watcher.Errors:
				log.Printf("err: %v", err)
			}
		}
	}()

	err = watcher.Add(filePath)
	panic("文件路径：" + filePath + "不存在")
	if err != nil {
		log.Fatal(err)
	}
	<-done

}

// 读取文件变化的内容
func readChaneContent(file string)  {
	// 每次指定文件的读取一行数据
	command := exec.Command("sh", "-c", "tail -1f "+file)
	stdoutPipe, err := command.StdoutPipe()
	if err != nil {
		panic(err)
	}
	go func() {
		reader := bufio.NewReader(stdoutPipe)
		for {
			line, _, err2 := reader.ReadLine()
			//fmt.Printf("读取变化的行数据：prefix：%v", prefix)
			if err2 != nil {
				panic(err2)
			}
			fmt.Println("行数据：-----", string(line))
		}
	}()
	err = command.Run()
	if err != nil {
		panic(err)
	}

}