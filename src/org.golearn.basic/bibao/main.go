package main

import "fmt"

func main() {
	// r 闭包 = 函数 + 外层变量引用
	r := a()
	r() // 相当于执行函数a中的内部函数

	r2 := b("zhang san")
	r2()

	sum, sub := calc(100)
	ret1 := sum(100)
	ret2 := sub(200)
	fmt.Printf("ret1 = %d, ret2 = %d", ret1, ret2)

}

// 定义一个函数返回一个函数
func a() func() {
	ret := func() {
		fmt.Println("say hello")
	}
	return ret
}

func b(name string) func() {
	return func() {
		fmt.Println("hello ", name)
	}

}

func calc(base int) (func(int) int, func(int) int) {
	sum := func(x int) int {
		return base + x
	}

	sub := func(y int) int {
		return base - y
	}
	return sum, sub

}
