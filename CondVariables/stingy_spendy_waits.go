package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	Money        = 100
	lock         sync.Mutex
	moneyDeposit = sync.NewCond(&lock)
)

func Stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		Money += 10
		fmt.Println("Stingy Sees Balance: ", Money)
		moneyDeposit.Signal()
		lock.Unlock()
		time.Sleep(time.Millisecond * 1)
	}
	fmt.Println("Stingy Done!")
}

func Spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		//账户余额不足20元时等待
		for Money < 20 {
			moneyDeposit.Wait()
		}
		Money -= 20
		fmt.Println("Spendy Sees Balance: ", Money)
		lock.Unlock()
		time.Sleep(time.Millisecond * 1)
	}
	fmt.Println("Spendy Done!")
}

func main() {
	go Stingy()
	go Spendy()
	time.Sleep(time.Millisecond * 3000)
	fmt.Println(Money)
}
