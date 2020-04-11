package main

import (
	"fmt"
)

// 函数
// fun name(入参) (返回参数)
func main() {

	data := sum(1, 2)
	fmt.Println(data)
	data2 := intSum2(1, 2, 4, 5, 6)
	fmt.Printf("可变参数：%d \n", data2)
	data3 := intSum3(1, 2, 4, 5,6,7,8)
	fmt.Printf("固定参数和可变参数使用：%d \n", data3)
	// 多个返回值的函数在接收返回值需要跟函数返回值一致
	data4, data5 := calc(2, 4)
	fmt.Printf("多个返回值：%d, %d \n", data4, data5)
	// 命名返回值
	data6, data7 := calc2(1, 2)
	fmt.Printf("多个返回值：data6 = %d, data7 = %d \n", data6, data7)

}

// 普通函数
func sum(x, y int) int {
	return x + y
}

// 可变参数函数
func intSum2(x ...int) int {
	count := 0;
	for i := 0; i < len(x); i++ {
		count += x[i]
	}
	return count
}

// 固定参数和普通参数混合使用
func intSum3(x int, y ...int) (int) {
	for i := 0; i < len(y); i++ {
		x += y[i]
	}

	return x
}

// 对个返回值
func calc(x int, y int) (int, int) {
	sum := x + y
	muti := x * y
	return sum, muti
}

// 命名返回值，就不需要在return后面添加相关参数了
func calc2(x int, y int) (sum , sub int)  {
	sum = x + y
	sub = x - y
	return
}