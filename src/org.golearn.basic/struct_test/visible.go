package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段的可见性

// 如果一个go语言结构体大写就表示可以公开访问（public）
// 使用结构体标签来对序列化后的字段做处理（驼峰命名）
type Student struct {
	ID int8 `json:"id"` // 指定序列化显示的字段名称
	Name string `json:"name"`
}

func  newStudent(ID int8, Name string) *Student {
	return &Student{
		ID:   ID,
		Name: Name,
	}

}

type class struct {
	Title string `json:"title"`
	Students []Student `json:"studentList"`
}

func main()  {
	c1 := class{
		Title:    "火箭101",
		Students: make([]Student, 0, 20),
	}

	// 像C1中添加学生
	for i := 0; i < 10 ; i++ {
		stu := newStudent(int8(i), fmt.Sprintf("stu %d", i))
		c1.Students = append(c1.Students, *stu)
	}
	fmt.Printf("class c1 %#v\n", c1)

	// json序列化
	data, err := json.Marshal(c1)
	if nil != err {
		fmt.Println("json 序列化失败， err:", err)
		return
	}
	fmt.Printf("序列化后的类型：data = %T\n", data) //  []uint8
	fmt.Printf("按字符串格式输出：data = %s\n", data)

	// json反序列化 json字符串转go语言数据
	//jsonStr := `{"Title":"火箭101","Students":[{"ID":0,"Name":"s":1,"Name":"stu 1"},{"ID":2,"Name":"stu 2"},{"ID":3,"Name":"stu 3"},{"ID":4,"Name":"stu 4"},{"ID":5,"Name":"stu 5"},{"ID":6,"Name":"stu 6"},{"ID":7,"Name":"stu 7"},{"ID":8,"Name":"stu 8"},{"ID":9,"Name":"stu 9"}]}`
	jsonStr := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`

	fmt.Println("=====反序列化====")
	// 反序列化对象要求是个指针类型
	var c2 *class = &class{}
	err = json.Unmarshal([]byte(jsonStr), c2)
	if nil != err {
		fmt.Println("json 反序列化失败， err:", err)
		return
	}

	fmt.Printf("反序列化后的类型：data = %T\n", c2) //  []uint8
	fmt.Printf("反序列化后的类型：data = %#v\n", c2) //  []uint8


}
