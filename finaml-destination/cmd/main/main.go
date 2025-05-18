package main

import (
	"Paper/finaml-destination/finaml-destination/internal/constants"
	"Paper/finaml-destination/finaml-destination/internal/usecase/benchmark"
	"Paper/finaml-destination/finaml-destination/internal/usecase/methods/polenomial"
	"fmt"
)

func main() {
	fmt.Printf("\nStart debugging!")
	benchmark.Benchmark(1000, constants.MatrixSize, polenomial.MinCost)
}
