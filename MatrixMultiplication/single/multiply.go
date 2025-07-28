package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MatrixSize = 250

var (
	matrixA [MatrixSize][MatrixSize]int
	matrixB [MatrixSize][MatrixSize]int
	result  [MatrixSize][MatrixSize]int
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
	for col := 0; col < MatrixSize; col++ {
		for i := 0; i < MatrixSize; i++ {
			result[row][col] += matrixA[row][i] * matrixB[i][col]
		}
	}
}
func main() {
	fmt.Println("Working......")
	now := time.Now()
	for i := 0; i < 100; i++ {
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		for row := 0; row < MatrixSize; row++ {
			workOutRow(row)
			//fmt.Println(result[row])
		}
	}
	fmt.Println("Done!")
	fmt.Printf("Time Taken: %v\n", time.Since(now))
}
