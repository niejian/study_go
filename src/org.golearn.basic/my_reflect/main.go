package main

import (
	"fmt"
	"reflect"
)

// 反射

// 声明反射
func reflectType(x interface{})  {
	// 不知道别人调用传递是什么类型的变量
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name())
}

type Cat struct {

}

type Dog struct {

}

func reflectVal(x interface{})  {
	val := reflect.ValueOf(x)
	// 打印值 类型
	fmt.Printf("%v, %T \n", val, val) // 1.23, reflect.Value, 并没有获取到val的真正值类型
	kind := val.Kind() // 拿到值对应的类型种类
	switch kind {
	case reflect.Float32:
		ret := float32(val.Float())
		fmt.Printf("ret  %v, %T \n", ret, ret)
	case reflect.Int32:
		ret := int32(val.Int())
		fmt.Printf("ret  %v, %T \n", ret, ret)

	}
}

func reflectSetVal(x interface{})  {
	val := reflect.ValueOf(x)
	// Elem()根据指针获取对应的地址值
	kind := val.Elem().Kind()
	switch kind {
	case reflect.Float32:
		val.Elem().SetFloat(3.14)
	case reflect.Int8:
		val.Elem().SetInt(100)

	}
}

func main()  {
	var a float32 = 1.23
	reflectType(a)
	var b int8 = 10
	reflectType(b)
	var c Cat
	reflectType(c) // main.Cat Cat
	var d Dog
	reflectType(d)
	fmt.Println("====reflect.ValueOf===")
	reflectVal(a)
	var e int8 = 2
	reflectSetVal(&e)
	fmt.Println(e)

}
