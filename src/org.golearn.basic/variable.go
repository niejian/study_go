package main

/**
  Constants:    true  false  iota  nil

Types:    int  int8  int16  int32  int64
		uint  uint8  uint16  uint32  uint64  uintptr
		float32  float64  complex128  complex64
		bool  byte  rune  string  error

Functions:   make  len  cap  new  append  copy  close  delete
		complex  real  imag
		panic  recover
*/
import (
	"fmt"
)

func foo() (int, string) {
	return 10, "张三"
}

func main() {
	fmt.Print("hello world")
	// 变量声明
	// var 变量名 变量类型
	var name string = "张三"
	fmt.Println(name)
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Print(a, b, c, d)
	// 一次初始化多个变量
	var name2, age2 = "zhangsan", 10
	// 类型推导，根据值的类型自动推导变量类型
	var name3 = "张三"
	var age3 = 10
	fmt.Printf(name2, age2, name3, age3)
	// 短变量
	m := 10
	n := 20
	fmt.Println(m, n)

	// 匿名变量
	fmt.Println("匿名变量 _")
	x, _ := foo()
	_, y := foo()
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	/*
		匿名变量不占用命名空间，不会分配内存，
		所以匿名变量之间不存在重复声明。 (在Lua等编程语言里，匿名变量也被叫做哑元变量。)
	*/

	fmt.Println("===常量===")
	const pi = 3.1415
	const e = 2.7182
	// 联合声明
	const (
		// 声明多个变量时省略了值则跟上一个值保持一致
		n1 = 100
		n2
		n3
	)

	fmt.Printf("pi %s e：%s, n1 %s,n2 %s,n3 %s", pi, e, n1, n2, n3)
	// iota 关键字
	/*
		iota在const关键字出现时被重置为0，const没增加一行将是iota计数+1
		可以理解为const关键字行索引。
	*/
	fmt.Println("\n====iota====")
	const (
		n4 = iota // 0
		n5 = iota // 1
		n6        // 2
		n7 = 100
		n8 = iota // 4
	)

	fmt.Println("n4, n5, n6, n8", n4, n5, n6, n8)
}
