package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

const (
	TotalAccounts  = 50000
	MaxAccountMove = 10
	InitialMoney   = 100
	TotalThreads   = 4
)

// PerformMove 模拟账户移动
/**
 * 模拟账户移动
 * @param totalLedger   账户列表
 * @param locks         账户锁
 * @param totalTransactions   总交易次数
 */
func PerformMove(totalLedger *[TotalAccounts]int32, locks *[TotalAccounts]sync.Locker, totalTransactions *int32) {
	for {
		accountA := rand.Intn(TotalAccounts)
		accountB := rand.Intn(TotalAccounts)
		for accountB == accountA {
			accountB = rand.Intn(TotalAccounts)
		}
		toLock := []int{accountA, accountB}
		amount := rand.Int31n(MaxAccountMove)
		sort.Ints(toLock)
		//按顺序上锁 不按顺序会导致死锁
		//goroutine 1 锁定账户 A 然后尝试锁定账户 B
		//同时 goroutine 2 锁定账户 B 然后尝试锁定账户 A
		//结果两个 goroutine 互相等待形成死锁
		locks[toLock[0]].Lock()
		locks[toLock[1]].Lock()
		atomic.AddInt32(&totalLedger[accountA], -amount)
		atomic.AddInt32(&totalLedger[accountB], amount)
		atomic.AddInt32(totalTransactions, 1)
		locks[toLock[1]].Unlock()
		locks[toLock[0]].Unlock()
	}
}

func main() {
	fmt.Println("Starting.............")
	var totalLedger [TotalAccounts]int32
	var locks [TotalAccounts]sync.Locker
	var totalTransactions int32

	for i := 0; i < TotalAccounts; i++ {
		totalLedger[i] = InitialMoney
		locks[i] = NewSpinLock()
	}

	for i := 0; i < TotalThreads; i++ {
		go PerformMove(&totalLedger, &locks, &totalTransactions)
	}
	for {
		time.Sleep(time.Second * 2)
		var sum int32
		for i := 0; i < TotalAccounts; i++ {
			locks[i].Lock()
		}
		for i := 0; i < TotalAccounts; i++ {
			sum += totalLedger[i]
		}
		for i := 0; i < TotalAccounts; i++ {
			locks[i].Unlock()
		}
		fmt.Printf("Total Transactions: %d, Total Money: %d\n", totalTransactions, sum)
	}
}
