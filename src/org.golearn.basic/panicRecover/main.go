package main

import "fmt"

// 使用panic、recover来处理异常
// recover只在defer调用的函数中有效
func main()  {
	A()
	B()
	C()
}

func A()  {
	fmt.Println("func A")
}

func B()  {
	// defer在panic之前定义
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()

	panic("panic in B")

}

func C()  {
	fmt.Println("func C")

}

