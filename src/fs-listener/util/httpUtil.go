package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	if "" == contentType {
		contentType = DEFAULT_CONTENT_TYPE
	}

	// 将结构体转换为json
	bytesData, err := json.MarshalIndent(data, "", "")
	if err != nil {
		panic("post json data 失败")
	}

	//fmt.Printf("链接：%v，请求参数：%v \n", url, string(bytesData))

	reader := bytes.NewReader(bytesData)
	response, err := http.Post(url, DEFAULT_CONTENT_TYPE, reader)

	if err != nil {
		panic(err)
	}
	fmt.Printf("  返回状态码：%v \n", response.Status)
	readBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&readBytes))

	return *str
}
