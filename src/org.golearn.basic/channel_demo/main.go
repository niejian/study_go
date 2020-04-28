package main

import "fmt"

// channel 多个goroutine直接的数据传递的通道

func main() {
	// 定义一个通道(channel是一种引用类型需要初始化才能使用)
	// ch := make(chan int) // 无缓冲区通道（会报思索）
	ch := make(chan int, 2) // 有缓冲区通道
	// 将一个值发送到channel中
	ch <- 10 //
	// 取值
	x := <- ch
	fmt.Println(x)
	//close(ch)
}
