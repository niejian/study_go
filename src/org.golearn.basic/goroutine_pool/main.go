package main

import (
	"fmt"
	"time"
)

/*
goroutine 连接池，
*/

func worker(index int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker: %d, start job: %d\n", index, job)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", index, job)
		// 向result中设值
		results <- job * 2
	}
}

// channel select
func channelSelect(ch chan int)  {
	for i := 0; i < 10; i++ {
		select {
		case x := <- ch: // channel中能拿的东西
			fmt.Println("channel 有值了",x)
		case ch <- i: // 给channel天健元素
			

		}
	}
}

func main() {

	// 连接池
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}
	// 发送任务
	for i := 0; i < 5; i++ {
		jobs <- i

	}
	close(jobs)
	// 输出结果
	for i := 0; i < 5; i++ {
		<-results
	}
}
