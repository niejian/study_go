package main

import "fmt"

// go 结构体，类型
// 类型定义
type myInt int // 通过type关键字定义，myInt是一种新的数据类型，具有 int的特点

//类型别名（小名，阿猫阿狗）
type NewInt = int

func main()  {
	// 类型定义和类型别名的区别
	var a myInt
	var b NewInt
	fmt.Printf("类型定义 a %T \n", a) // 类型定义 a main.myInt.表示的是main包下的myInt类型
	fmt.Printf("类型别名  %T \n", b) // 类型别名  int

	

}
