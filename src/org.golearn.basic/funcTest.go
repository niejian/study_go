package main

import (
	"errors"
	"fmt"
	"strings"
)

// 函数
// fun name(入参) (返回参数)
// 变量作用域

// 如果全局变量和局部变量重名，优先访问局部变量
var m int = 10


var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
)

func main() {

	data := sum(1, 2)
	fmt.Println(data)
	data2 := intSum2(1, 2, 4, 5, 6)
	fmt.Printf("可变参数：%d \n", data2)
	data3 := intSum3(1, 2, 4, 5, 6, 7, 8)
	fmt.Printf("固定参数和可变参数使用：%d \n", data3)
	// 多个返回值的函数在接收返回值需要跟函数返回值一致
	data4, data5 := calc(2, 4)
	fmt.Printf("多个返回值：%d, %d \n", data4, data5)
	// 命名返回值
	data6, data7 := calc2(1, 2)
	fmt.Printf("多个返回值：data6 = %d, data7 = %d \n", data6, data7)
	m := -1 // 与全局变量m重名
	fmt.Printf("访问重名的全局变量与局部变量，优先返回局部变量 %d \n", m)
	fmt.Println("=====类型函数===")
	// 函数类型与变量，用来定义特定入参和出参的函数，
	// type calcution(int, int) (int) 只要满足两个int类型入参，返回值是int的函数都是calcution类型函数
	type calculation func(int, int) int
	var funType calculation
	funType = sum
	fmt.Printf("funType类型: %T \n", funType)
	fmt.Println(funType(1, 2))
	f := calc
	fmt.Printf("f 类型： %T \n", f)
	fmt.Println(f(1, 2))

	fmt.Println("===高阶函数-函数作为参数===")
	// 将函数作为参数来声明一个函数
	ret2 := advanceCalc(10, 1, sum)
	fmt.Println("函数作为参数返回：", ret2)
	fmt.Print("----函数作为返回值----")
	a, b := do("+")
	fmt.Printf("a %T, b %T", a, b)
	fmt.Println(a, b)
	// 匿名函数，定义在函数里面的函数
	any := func(x, y int) {
		fmt.Println(x + y)

	}
	fmt.Println("==匿名函数：定义在函数里面的函数==")
	any(10, 20)

	// 函数闭包
	// 函数的闭包 = 函数 + 外层变量的引用
	fmt.Println("===函数的闭包==")
	r := test()
	r()
	r2 := test2("李思")
	r2()

	r3 := makeSuffixFunc(".avi")
	// 上山  传入到返回函数作为入参
	fmt.Println(r3("上山"))
	f2, f3 := advanceCalc2(10)
	fmt.Println(f2(1)) // 10 + 1
	fmt.Println(f3(2)) // 10 + 1 - 2

	dataMap := dispatchCoin(users, coins)
	if nil != dataMap && len(dataMap) > 0 {
		for username, count := range dataMap {
			fmt.Println(username , " get ", count , " coins")
		}
	}


}

// 普通函数
func sum(x, y int) int {
	return x + y
}

// 可变参数函数
func intSum2(x ...int) int {
	count := 0
	for i := 0; i < len(x); i++ {
		count += x[i]
	}
	return count
}

// 固定参数和普通参数混合使用
func intSum3(x int, y ...int) int {
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
func calc2(x int, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}
func sub(x, y int) int {
	return x - y

}

// 将函数作为参数来声明一个函数
func advanceCalc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值 (两个入参， 出参有两个，一个int类型，一个错误类型)
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return sum, nil
	case "-":
		return sub, nil
	default:
		//err := "无法识别的操作符"
		return nil, errors.New("无法识别的操作符")

	}

}

// 函数的闭包 = 函数 + 外层变量的引用
// 定义一个函数，返回值是一个函数
func test() func() {
	name := "zhangsan"
	return func() {
		// 执行匿名函数时，本函数没有找到变量name，找到外层的变量声明
		fmt.Println("hello，", name)
	}
}

func test2(name string) func() {
	return func() {
		// 执行匿名函数时，本函数没有找到变量name，找到外层的变量声明
		fmt.Println("hello，", name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

func advanceCalc2(base int) (func(int) int, func(int) int) {
	sum1 := func(i int) int {
		base += i
		return base
	}
	sub2 := func(i int) int {
		base -= i
		return base
	}

	return sum1, sub2
}

func dispatchCoin(users []string, total int ) map[string]int {
	datamap := make(map[string]int, 8)

	for _, username := range users {
		count := 0
		if strings.Contains(username,"e") || strings.Contains(username, "e") {
			count += strings.Count(username, "e")
			count += strings.Count(username, "E")
		}

		if strings.Contains(username,"i") || strings.Contains(username, "I") {
			count += strings.Count(username, "i") * 2
			count += strings.Count(username, "I") * 2
		}

		if strings.Contains(username,"o") || strings.Contains(username, "O") {
			count += strings.Count(username, "o") * 3
			count += strings.Count(username, "O") * 3
		}

		if strings.Contains(username,"U") || strings.Contains(username, "u") {
			count += strings.Count(username, "u") * 4
			count += strings.Count(username, "U") * 4
		}
		datamap[username] = count

	}

	return datamap
}
