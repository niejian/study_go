package main

import (
	"fmt"
	"sync"
)

// go协程 goroutine

/*
goroutine的调度（GPM）
G: 一个个的goroutine
P: process(机器的处理核心线程数)
M: mgo运行时的mechine
P和M是一一对应的；P管理着一组G挂载在M上运行，当一个G长久的阻塞在M上，runtime会创建一个新的M
阻塞G所在的P会把其他的G挂载在新的M上，当旧的G阻塞完成后，回收旧的M
*/

// 相当于java的countlatch
var wg sync.WaitGroup

func hello() {
	fmt.Println("hello..world")
	// countlatch - 1
	wg.Done()
}

func main() {

	wg.Add(1000)
	// 主线程开启goroutine去执行hello函数
	go hello()

	for i := 1; i <= 1000 ; i++  {
		go func(i int) {
			// 如果直接使用i,那么i的输出就总是循环到最后的那个数
			fmt.Println("匿名函数:", i)
			wg.Done()
		}(i)
	}

	fmt.Println("hello main")



	// 可能只执行main函数中的打印语句
	// countdownlatch = 1
	wg.Wait()
}
