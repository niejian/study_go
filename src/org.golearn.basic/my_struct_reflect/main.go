package main

import (
	"fmt"
	"reflect"
)

// 结构体反射
type student struct {
	Name string `json:"name" ini:"s_name"`
	Score int `json:"score" ini:"s_score"`
}

func (s student) Study() string  {
	msg := "好好学习，天天向上"
	fmt.Println(msg)
	return msg

}

func (s student) Sleep() string  {
	var msg string = "hello"
	fmt.Println(msg)
	return msg

}

func printMethod(x interface{})  {

	// 获取类型
	t := reflect.TypeOf(x)
	// 值类型
	v := reflect.ValueOf(x)
	// 获取结构体中的方法数量
	numMethod := t.NumMethod()
	fmt.Printf("参数：%s 方法数量：%v\n", t, numMethod)

	for i := 0; i < numMethod ; i++  {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)

	}

}

func main()  {
	stu1 := student{
		Name:  "zhang san",
		Score: 80,
	}

	val := reflect.TypeOf(stu1)
	fmt.Printf("name:%v, kind:%v\n", val.Name(), val.Kind())
	// 获取结构体变量所有字段
	for i := 0; i < val.NumField(); i++  {
		fieldObj := val.Field(i)
		fmt.Printf("name:%v, type:%v, tag:%v\n", fieldObj.Name, fieldObj.Type, fieldObj.Tag)
		// 获取tag中对应的值
		fmt.Println( fieldObj.Tag.Get("json"), fieldObj.Tag.Get("ini"))

	}
	printMethod(stu1)


}
