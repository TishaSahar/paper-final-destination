package benchmark

import (
	"fmt"
	"math/rand/v2"
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

// PrintMatrix prints the matrix in a readable format.
func PrintMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
