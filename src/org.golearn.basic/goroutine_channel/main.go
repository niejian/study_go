package main

import "fmt"

/**
两个goroutine两个channel
 1. 生成0-100的数字发送到cha1
 2. 从chan1取出数据计算平方，发送到chan2
 */
// 生成数据
func Generate(ch chan<- int)  {
	for i := 0; i < 100 ; i++  {
		// 单项通道
		ch <- i
	}
	close(ch)
}

// 设置单向通道
func getData(ch1 <-chan int, cha2 chan<- int) {
	for  {
		tmp, ok := <- ch1
		if !ok {
			break
		}
		cha2 <- tmp * tmp
	}
	close(cha2)
}


func main()  {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go Generate(ch1)
	go getData(ch1, ch2)

	for ret := range ch2 {
		fmt.Println(ret)
	}

}
