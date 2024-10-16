package main

import "fmt"

// import "fmt"

func main() {
	matrix1 := [][]int {
		{11, 12, 13},
		{21, 22, 23},
	}

	matrix2 := [][] int {
		{1, 2, 3},
		{4, 5, 6},
	}
	
	totalMatrix, err:= AddMatrix(matrix1, matrix2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	PrintMatrix(totalMatrix, "Sum of Matrix")
}