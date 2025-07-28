package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const MatrixSize = 250

var (
	matrixA  [MatrixSize][MatrixSize]int
	matrixB  [MatrixSize][MatrixSize]int
	result   [MatrixSize][MatrixSize]int
	rwLock   sync.RWMutex                     // 读写锁，保护矩阵访问
	wg       sync.WaitGroup                   // 等待组，同步计算完成
	condLock = sync.NewCond(rwLock.RLocker()) // 条件变量，用于线程间通信
)

// generateRandomMatrix 随机生成矩阵区间在 [-5,5)
func generateRandomMatrix(matrix *[MatrixSize][MatrixSize]int) {
	for row := 0; row < MatrixSize; row++ {
		for col := 0; col < MatrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}
func workOutRow(row int) {
	//上读锁保证矩阵不会被修改
	rwLock.RLock()
	/**
	无限循环是为了让 goroutine 在 100 轮计算中重复使用而不是每次新建 goroutine
	*/
	for {
		for col := 0; col < MatrixSize; col++ {
			for i := 0; i < MatrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
		wg.Done()       //上一个线程已经完成计算
		condLock.Wait() //解锁等待下一个线程开始计算
	}
}
func main() {
	fmt.Println("Working......")
	//启动多个goroutine
	wg.Add(MatrixSize)
	for row := 0; row < MatrixSize; row++ {
		go func(int) {
			workOutRow(row)
		}(row)
		//fmt.Println(result[row])
	}
	now := time.Now()
	for i := 0; i < 100; i++ {
		//生成新的矩阵前需要确保矩阵计算完成
		wg.Wait()
		/**
		写锁的作用是保证数据一致性：
		当主线程需要更新 matrixA 和 matrixB 时必须确保没有其他 goroutine 正在读取这些矩阵
		写锁会阻止所有读操作确保在生成新矩阵的过程中工作 goroutine 不会读取到不一致的数据
		*/
		rwLock.Lock()
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		rwLock.Unlock()
		//生成结束表示已经完成一轮计算 此时等待组进行重新初始化
		wg.Add(MatrixSize)
		//条件锁广播唤醒所有等待进程
		condLock.Broadcast()
	}
	fmt.Println("Done!")
	fmt.Printf("Time Taken: %v\n", time.Since(now))
}
