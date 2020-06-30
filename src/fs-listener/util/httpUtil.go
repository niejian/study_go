package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"unsafe"
)

const (
	DEFAULT_CONTENT_TYPE = "application/json;charset=UTF-8"
	POST = "post"
	CONTENT_TYPE_KEY = "Content-Type"
)

func Get(url, param string)  string {
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		// 读取长度
		readLen, err := response.Body.Read(buffer[0:])
		result.Write(buffer[0: readLen])
		if err != nil && err == io.EOF {
			break
		}else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

func Post(url string, data interface{}, contentType string) string {

	// 异常处理
	defer func() {
		if r := recover(); r!=nil{
			log.Printf("post err: %v",r)
		}
	}()

	if "" == contentType {
		contentType = DEFAULT_CONTENT_TYPE
	}


	var response = doPost(url, data, "")
	return response
}
//defer 异常处理，发生异常，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点
func doPost(url string, data interface{}, contentType string) string {

	defer func() {
		if r := recover(); r!=nil {
			log.Printf("post 请求发生错误: %v",r)
		}
	}()

	// 将结构体转换为json
	bytesData, _ := json.MarshalIndent(data, "", "")


	//fmt.Printf("链接：%v，请求参数：%v \n", url, string(bytesData))

	reader := bytes.NewReader(bytesData)

	response, err := http.Post(url, DEFAULT_CONTENT_TYPE, reader)
	if err != nil {
		log.Printf("post %v, 请求失败, 重试，%v\n", url, err)
		//time.Sleep(1 * time.Second)
		//Post(url, data, "")
		panic(err)
	}
	//fmt.Printf("  返回状态码：%v \n", response.Status)
	readBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("解析返回结果失败：%v \n", err)
		//panic(err)
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&readBytes))
	log.Println("请求结果信息：", *str)

	return *str

}
