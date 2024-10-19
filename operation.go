package main

import (
	"fmt"
)

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

// create subtract
func SubtractMatrix(matrix1, matrix2 [][]int) ([][]int, error) {
	if len(matrix1[0]) != len(matrix2[0]) || len(matrix1) != len(matrix2) {
		return nil, fmt.Errorf("matrix must be in the same dimensions")
	}
	cols := len(matrix1[0])
	rows := len(matrix1)

	sumMatrix := CreateZeroMatrix(rows, cols)
	fmt.Println(sumMatrix)
	
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sumMatrix[row][col] = matrix1[row][col] - matrix2[row][col]
		}
	}
	return sumMatrix, nil
}

func DotProduct(matrix1, matrix2 [][] int) ([][] int, error) {
	row1 := len(matrix1)
	col1 := len(matrix1[0])

	row2 := len(matrix2)
	col2 := len(matrix2[0])

	if col1 != row2 {
		return nil, fmt.Errorf("matrix dimensions cannot be multiplied! Wrong dimension: %dx%d & %dx%d", row1, col1, row2, col2)
	}

	prodMatrix := CreateZeroMatrix(row1, col2)
	
	// Loop through prodMatrix to assign value
	for row := 0; row < len(prodMatrix); row++ {
		for col := 0; col < len(prodMatrix[0]); col++ {
			
			for col1row2 := 0; col1row2 < col1; col1row2++ {
				prodMatrix[row][col] += matrix1[row][col1row2] * matrix2[col1row2][col]
			}
		}
	}


	return prodMatrix, nil
}
// create transpose
func TransposeMatrix(matrix [][]int) [][]int {

	rows := len(matrix[0])
	cols := len(matrix)

	tMatrix := CreateZeroMatrix(rows, cols)
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			tMatrix[row][col] = matrix[col][row]
		}
	}
	
	
	return tMatrix
}
func CreateZeroMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}
func PrintMatrix(matrixName string, matrix [][]int) {
	fmt.Println()
	fmt.Println(matrixName, ": ", len(matrix), "x", len(matrix[0]))
	for _, row := range matrix {
		fmt.Println(row);
	}
}