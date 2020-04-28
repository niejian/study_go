package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine的锁和安全

// 多个goroutine操作一个全局变量
var(
	x int
	wg sync.WaitGroup
	lock sync.Mutex // 互斥锁，同一时刻，只有一个goroutine获取到锁
	rwLock sync.RWMutex // 读写互斥锁
)

func add()  {
	for i := 0; i<5000 ; i++  {
		// 加锁，使操作安全
		lock.Lock()
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}

// 读写互斥锁demo

func read()  {
	//lock.Lock()
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwLock.RUnlock()
	wg.Done()

}

func write()  {
	//lock.Lock()
	rwLock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 10)
	//lock.Unlock()
	rwLock.Unlock()
	wg.Done()

}

func main()  {
	start := time.Now()
	//wg.Add( 2)
	//go add()
	//go add()
	// 读写互斥锁
	for i := 0; i< 10 ; i++  {
		wg.Add(1)
		go write()
	}
	for i := 0; i< 5000 ; i++  {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	//fmt.Println(x)
	fmt.Println("花费时间：", time.Now().Sub(start))

}
