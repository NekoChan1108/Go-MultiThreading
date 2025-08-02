package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	//Money = 100
	//lock  sync.Mutex
	Money int64 = 100
)

func Stingy() {
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		//Money += 10
		//lock.Unlock()
		//原子操作
		atomic.AddInt64(&Money, 10)
		time.Sleep(time.Millisecond * 1)
	}
}

func Spendy() {
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		//Money -= 10
		//lock.Unlock()
		//原子操作
		atomic.AddInt64(&Money, -10)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	go Stingy()
	go Spendy()
	time.Sleep(time.Millisecond * 3000)
	fmt.Println(Money)
}
