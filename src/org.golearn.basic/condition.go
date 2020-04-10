package main

import (
	"fmt"
)

// 循环判断
func main() {

	demo4()
	fmt.Println("switch case...")
	result := switchCase(2)
	fmt.Println(result)
	switchdemo2()
	fmt.Println("not goto....")
	notGotoDemo()
	fmt.Println("not goto....")

	gotoDemo()
	nineMutiNie(10)
}

// 有限循环
func demo1() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// 循环不设置初始条件
func demo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	for i < 10 {
		fmt.Println(i)
		i++
	}
}

// 无线循环
func demo3() {
	for {
		fmt.Println("无限循环...")
	}
	fmt.Println("五险")
}

// for range
/*
	通过 for range遍历数组、切片、字符串、map以及通道（channel），返回值有以下规律：
	1. 数组、切片、字符串返回索引和值
	2. map返回键值
	3. 通道（channel）返回通道内的值

*/

func demo4() {
	fmt.Println("for range 遍历字符串")
	for _, r := range "hello web" {
		fmt.Println(r)
	}
}

func switchCase(finger int) string {
	result := "无效输入"
	switch finger {
	case 1:
		result = "大拇指"
	case 2:
		result = "食指"
	case 3:
		result = "中指"
	case 4:
		result = "无名指"
	case 5:
		result = "小拇指"
	default:
		result = "无效输入"
	}
	return result
}

// 一个分支有多个值判断
func mutiSwitch() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 0:
		fmt.Println("偶数")
	}
}

// fallthrough 可以使满足条件的下一个条件也执行
func switchdemo2() {
	s := "a"
	switch {
	case s == "A":
		fmt.Println("A")
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("default")

	}
}

// goto跳转到指定标签
func notGotoDemo() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("i = %d, j = %d \n", i, j)
		}

		if breakFlag {
			break
		}
	}
}

// 使用goto后
func gotoDemo() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 跳转到退出标签，并执行退出标签里的逻辑
				goto breakTag
			}
			fmt.Printf("i = %d, j = %d \n", i, j)
		}
	}
	return

breakTag:
	fmt.Println("结束for循环")

}

func nineMutiNie(x int) {
	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d", j, i, i*j)
			fmt.Print("\t")
		}
		fmt.Println()

	}
}
