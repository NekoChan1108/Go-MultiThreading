package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1, lock2 sync.Mutex
)

func blueRobot() {
	for {
		fmt.Println("Blue Acquire Lock1...")
		lock1.Lock()
		fmt.Println("Blue Acquire Lock2...")
		lock2.Lock()
		fmt.Println("Blue Acquire Both...")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Blue Release Both...")
	}
}
func redRobot() {
	for {
		fmt.Println("Red Acquire Lock2...")
		lock2.Lock()
		fmt.Println("Red Acquire Lock1...")
		lock1.Lock()
		fmt.Println("Red Acquire Both...")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Red Release Both...")
	}
}
func main() {
	go func() {
		blueRobot()
	}()
	go func() {
		redRobot()
	}()
	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}
