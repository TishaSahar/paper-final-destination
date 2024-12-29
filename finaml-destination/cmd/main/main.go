package main

import (
	"Paper/finaml-destination/finaml-destination/internal/constants"
	"Paper/finaml-destination/finaml-destination/internal/usecase/benchmark"
)

func main() {
	matrix := benchmark.GenerateMatrix(constants.MatrixSize)
	benchmark.PrintMatrix(matrix[0:5][0:5])
}
