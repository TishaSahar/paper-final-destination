package main

import (
	"Paper/finaml-destination/finaml-destination/internal/constants"
	"Paper/finaml-destination/finaml-destination/internal/usecase/benchmark"
	"Paper/finaml-destination/finaml-destination/internal/usecase/methods/polenomial"
)

func main() {
	// benchmark.Benchmark(30, constants.MatrixSize, dynamic.MinCostAssignment)
	benchmark.Benchmark(1000, constants.MatrixSize, polenomial.MinCost)
}
