package main

import (
	"fmt"
	"time"
)

var (
	Money = 100
	//lock  Sync.Mutex
)

func Stingy() {
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		Money += 10
		//lock.Unlock()
		time.Sleep(time.Millisecond * 1)
	}
}

func Spendy() {
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		Money -= 10
		//lock.Unlock()
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	go Stingy()
	go Spendy()
	time.Sleep(time.Millisecond * 3000)
	fmt.Println(Money)
}
