package main

import "fmt"

// go 语言中的异常处理机制（panic、recovery）
// panic: 产生异常，
func GO()  {
	fmt.Println("调用go函数，没发生异常，继续往下执行")
}
func Java()  {
	// 如果只是做了panic捕获，程序在执行到panic上就不会再往下执行了
	// 现在添加recovery 来处理
	defer func() {
		err := recover()
		if nil != err{
			fmt.Printf("捕获到panic产生的异常 err： %v \n", err)
			// recovery捕获到异常后，直接跳出这个方法，在主函数中继续往下执行
			fmt.Println("我是defer里的匿名函数，我捕获到panic的异常了，我要recover，恢复过来了。")

		}


	}()
	panic("Java方法，抛出了一个异常")
	fmt.Println("调用Java方法，发生异常panic后的执行情况")

}

func Py()  {
	maps := make(map[string]int32, 10)
	fmt.Println("调用python函数，没发生异常，继续往下执行")

}

func main()  {
	GO()
	Java()
	fmt.Printf("panic后的劫后重生")
	Py()
}
