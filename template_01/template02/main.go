package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Gender string
	Age string
}

func hello(w http.ResponseWriter, r *http.Request)  {
	// 解析指定文件生成的模板对象
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	u1 := User{
		Name:   "张三",
		Gender: "男",
		Age:    "16",
	}
	m := map[string]string{
		"Name":"a",
		"Gender":"a",
		"Age":"2",
	}

	// 传递slice数据
	hobbies := []string{"a", "b", "c", "d" }

	data := map[string]interface{} {
		"m":m,
		"u1":u1,
		"hobbies": hobbies,
	}

	tmpl.Execute(w, data)
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
