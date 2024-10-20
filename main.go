package main

import (
	"log"
	// "fmt"
)

func main() {
	weight := [][]float32 {
		{5, 2},
		{3, 6},

	}

	input := [][] float32 {
		{0.1},
		{0.2},
	} 	

	prod, err := ActFunction[float32](weight, input, 6)

	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	PrintMatrix[float32]("neuron: ", prod)
}