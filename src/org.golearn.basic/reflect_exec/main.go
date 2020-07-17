package main

import (
	"fmt"
	"reflect"
)

// 反射
/**
通过反射可以获得变量的 类型信息和值信息
*/

// 获取变量的类型信息
func getReflectType(in interface{})  {
	v := reflect.TypeOf(in)
	// 获取变量的type和kind（底层数据类型）
	fmt.Printf("type: %v, kind: %v \n", v.Name(), v.Kind())
}

/*
类型(Type)：使用type关键字定义的类型信息
种类(Kind)：go语言中的底层数据类型
*/

type myInt int64

type person struct {
	Name string `json:"name"`
	Age int8 `json:"age"`
}

type book struct {
	Name string
	Price float32
}

// 通过反射的方式获取对应的值信息
func reflectValue(x interface{})  {
	val := reflect.ValueOf(x)
	kind := val.Kind() // 值对应的底层数据类型
	switch kind {
	case reflect.Int32:
		fmt.Printf("type is int32, value is %v \n", int32(val.Int()))
	case reflect.Int64:
		fmt.Printf("type is int64, value is %v \n", int64(val.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %v \n", float32(val.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %v \n",val.Float())
	case reflect.String:
		fmt.Printf("type is string, value is %v \n", val.String())
	case reflect.Map:
		fmt.Printf("type is map, value is %v \n", val.String())
	case reflect.Slice:
		fmt.Printf("type is slice, value is %v \n", val.String())
	}
}

// 通过反射设置值
func reflectSetVal(x interface{})  {
	// 通过反射获取到原来的值信息
	value := reflect.ValueOf(x)
	// 判断value的类型信息（Elem()只识别接口、指针类型）
	kind := value.Elem().Kind()
	switch kind {
	case reflect.Float64:
		newVal := value.Elem().Float() + 1.0
		// 反射中使用 Elem()方法获取指针对应的值
		value.Elem().SetFloat(newVal)
	case reflect.Int64:
		newVal := value.Elem().Int() + 1
		// 反射中使用 Elem()方法获取指针对应的值
		value.Elem().SetInt(newVal)
	}
}

type student struct {
	Name string
	Age int8
}

// 结构体反射
func (s student) Study() string  {
	msg := "study..."
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string  {
	msg := "sleep..."
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{})  {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	// 方法必须是public
	fmt.Printf("type NumMethod size:  %v\n", t.NumMethod())
	fmt.Printf("value NumMethod size: %v\n", v.NumMethod())

	// 获取x的方法信息
	methodNum := v.NumMethod()
	for i := 0; i < methodNum; i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("方法名称:  %v \n", t.Method(i).Name)
		fmt.Printf("方法类型:  %v \n", methodType)
		// 通过反射的方式调用方法
		var args = []reflect.Value{}
		// invoke
		v.Method(i).Call(args)
	}
}


func main()  {
	a := 2.13
	b := make([]int, 10)
	c := make(map[string]string, 10)
	var d myInt
	var e rune // 类型别名
	var f *int // 指针类型
	getReflectType(a)
	getReflectType(b)
	getReflectType(c)
	getReflectType(d)  // type: myInt, kind: int64
	getReflectType(e) // type: int32, kind: int32
	getReflectType(f) // type: , kind: ptr

	// 判断结构体类型
	g := person{
		Name: "a",
		Age:  10,
	}
	h := book{
		Name:  "sds",
		Price: 0.1 ,
	}
	getReflectType(g) // type: person, kind: struct
	getReflectType(h) // type: book, kind: struct

	// 通过反射获取值信息
	var aa float64 = 3.13
	var bb int64 = 10
	var cc = make([]int, 10)
	cc[0] = 1
	var dd = make(map[string]string, 10)
	dd["a"] = "a"
	dd["c"] = "c"
	reflectValue(aa)
	reflectValue(bb)
	reflectValue(cc)
	reflectValue(dd)
	fmt.Println("通过反射设置值")
	// 注意此处必须传指针类型，不然会panic
	reflectSetVal(&aa)
	reflectSetVal(&bb)
	fmt.Printf("重新设置后：aa=%v,  bb=%v \n", aa, bb)

	// 结构体反射
	personReflect := reflect.TypeOf(g)
	fmt.Printf("结构体反射：结构体名称：%v, 底层类型：%v\n", personReflect.Name(), personReflect.Kind())
	// 遍历结构体的所有字段信息
	for i := 0; i < personReflect.NumField(); i++ {
		field := personReflect.Field(i) // 字段名称
		fmt.Printf("字段名称：%v, index: %d, type: %v, tag: %v \n", field.Name, field.Index, field.Type, field.Tag)
	}

	// 通过字段名获取指定结构体字段信息
	if field, ok := personReflect.FieldByName("Name"); ok {
		fmt.Printf("name: %s, index: %d, type: %v, json tag: %v \n", field.Name, field.Index, field.Type, field.Tag)
	}

	stu := student{
		Name: "stu",
		Age:  10,
	}
	printMethod(stu)
}
