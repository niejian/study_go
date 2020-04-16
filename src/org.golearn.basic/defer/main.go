package main

import (
	"fmt"
)

// defer: 延迟执行，当函数执行道最后的时候，defer定义的就会执行，并且是倒序执行
// defer的执行时机是返回值赋值之后，ret指令执行之前
func main() {
	// start end 3 2 1
	//fmt.Println("start...")
	//defer fmt.Println("1")
	//defer fmt.Println("2")
	//defer fmt.Println("3")
	//fmt.Println("end")
	//fmt.Println("=======")
	//// defer在返回值赋值之后，return指令执行之前镜像
	//fmt.Println("f1 返回值", f1()) //5
	//fmt.Println("f2 返回值", f2())
	//fmt.Println("f3 返回值", f3())
	//fmt.Println("f4 返回值", f4())
	//
	//fmt.Println("------------")

	x := 1
	y := 2
	defer cacl("AA", x, cacl("A", x, y)) // 后执行：x = 10
	// 重新赋值
	x = 10
	defer cacl("BB", x, cacl("B", x , y)) // 先执行； x = 10， y = 20
	// B 10 20
	// BB 10 30
	y = 20
}

/*
defer 经典案例
*/
func f1() int{
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	fmt.Println("f2() ==> x", x)
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func cacl(index string, a, b int) int{
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret

}

