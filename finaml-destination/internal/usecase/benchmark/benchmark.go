package benchmark

import (
	"fmt"
	"math/rand/v2"
	"sync/atomic"
	"time"
)

// GenerateMatrix generates an n x n matrix filled with integers starting from 1.
func GenerateMatrix(n int) [][]int {
	// Create a 2D slice with n rows
	matrix := make([][]int, n)

	// Fill the matrix with random values
	randomizer := rand.New(rand.NewPCG(uint64(time.Now().UnixMilli()), 0))
	for i := 0; i < n; i++ {
		// Initialize each row with n columns
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = randomizer.IntN(n-0) + 0
		}
	}

	return matrix
}

// Initialize MW vector
func GenerateMW(n int) []int {
	randomizer := rand.New(rand.NewPCG(uint64(time.Now().UnixMilli()), 0))
	vector := make([]int, n)
	for j := 0; j < n; j++ {
		vector[j] = randomizer.IntN(n-0) + 0
	}

	return vector
}

// PrintMatrix prints the matrix in a readable format.
func PrintMW(vetor []int) {
	for j := 0; j < len(vetor); j++ {
		fmt.Printf("%d ", vetor[j])
	}
}

// PrintMatrix prints the matrix in a readable format.
func PrintMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func BenchMethod(testSize int, method func(tree [][]int, MW []int) int) int64 {
	matrix := GenerateMatrix(testSize)
	mw := GenerateMW(testSize)
	// Calculate working time
	startTime := time.Now()
	method(matrix, mw)
	return time.Now().Local().Sub(startTime).Microseconds()
}

func Benchmark(iterationsCount, testSize int, method func(tree [][]int, MW []int) int) {
	var avg atomic.Int64
	for range iterationsCount {
		// append iteration working time
		avg.Swap(avg.Add(BenchMethod(testSize, method)))
	}
	fmt.Printf("\n=== Average working time of given method is %f microseconds\n", float64(avg.Load())/float64(iterationsCount))
}
