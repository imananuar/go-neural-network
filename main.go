package main

import "fmt"

// import "fmt"

// import "fmt"

func main() {
	matrix1 := [][]int {
		{1, 2, 3},
		{4, 5, 6},
		{4, 5, 6},

	}

	matrix2 := [][] int {
		{1},
		{4},
		{1},
	}

	// totalMatrix, err:= AddMatrix(matrix1, matrix2)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// tMatrix := TransposeMatrix(matrix1);

	// subtractMatrix, err:= SubtractMatrix(matrix1, matrix2)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// PrintMatrix("Sum of Matrix", totalMatrix)
	// PrintMatrix("Transpose Matrix", tMatrix)
	// PrintMatrix("Subtract Matrix", subtractMatrix)
	prod, err := DotProduct(matrix1, matrix2)
	if err != nil {
		fmt.Println("Dot Product Error: ", err)
	}
	PrintMatrix("Product Matrix", prod)
}