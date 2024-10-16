package main

import "fmt"

func AddMatrix(matrix1, matrix2 [][]int) ([][]int, error) {

	if len(matrix1[0]) != len(matrix2[0]) || len(matrix1) != len(matrix2) {
		return nil, fmt.Errorf("matrices must be in the same dimensions")
	}
	cols := len(matrix1[0])
	rows := len(matrix1)

	sumMatrix := CreateZeroMatrix(rows, cols)
	fmt.Println(sumMatrix)
	
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sumMatrix[row][col] = matrix1[row][col] + matrix2[row][col]
		}
	}
	return sumMatrix, nil
}

func CreateZeroMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

func PrintMatrix(matrix [][]int, matrixName string) {
	fmt.Println()
	fmt.Println(matrixName, ": ")
	for _, row := range matrix {
		fmt.Println(row);
	}
}