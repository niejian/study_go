package main

import (
	"fmt"
	"strings"
)

/**
数组：
*/
func main() {
	// 数组声明
	// var 变量名[数组长度] T
	var a [3]int
	for _, r := range a {
		// 未赋值的数组元素给默认值
		fmt.Println(r)
	}
	s := strings.Count("abc")
	fmt.Println(s)
	var b = [3]int{1, 3}
	fmt.Println(b)
	var c [3]string
	fmt.Println(c)
	// 不定长的数组长度赋值
	var d = [...]int{1, 2, 3}
	fmt.Println(d)
	fmt.Printf("array d: %T", d) // array d: [3]int%

	// 给指定索引的值赋值，下标1=12，3 = 1
	e := [...]int{1: 12, 3, 3: 1, 9}
	fmt.Println(e)

	// 数组变量
	for index, r := range e {
		println(index, r)
	}

	// 多维数组
	f := [3][2]string{
		{"1", "2"},
		{"11", "22"},
		{"111", "222"},
	}
	fmt.Printf(f[0][1])

}
