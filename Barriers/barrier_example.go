package main

import (
	"fmt"
	"time"
)

func waitOnBarrier(name string, duration int, barrier *Barrier) {
	for {
		fmt.Println(name, "Running")
		time.Sleep(time.Duration(duration) * time.Second)
		fmt.Println(name, "Waiting")
		barrier.Wait()
	}
}
func main() {
	barrier := NewBarrier(2)
	go waitOnBarrier("Red", 4, barrier)
	go waitOnBarrier("Blue", 10, barrier)
	time.Sleep(100 * time.Second)
}
