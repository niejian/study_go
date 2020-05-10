package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	// 解析指定文件生成的模板对象
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	tmpl.Execute(w, time.Now())
}
//template模板引擎
func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("服务启动失败，%v\n", err)
		return
	}



}
