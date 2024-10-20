package main

import (
	"fmt"
)

func AddMatrix[T Matrix](matrix1, matrix2 [][]T) ([][]T, error) {

	if len(matrix1[0]) != len(matrix2[0]) || len(matrix1) != len(matrix2) {
		return nil, fmt.Errorf("matrix must be in the same dimensions")
	}
	cols := len(matrix1[0])
	rows := len(matrix1)

	sumMatrix := CreateZeroMatrix[T](rows, cols)
	
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sumMatrix[row][col] = matrix1[row][col] + matrix2[row][col]
		}
	}
	return sumMatrix, nil
}
func SubtractMatrix[T Matrix](matrix1, matrix2 [][]T) ([][]T, error) {
	if len(matrix1[0]) != len(matrix2[0]) || len(matrix1) != len(matrix2) {
		return nil, fmt.Errorf("matrix must be in the same dimensions")
	}
	cols := len(matrix1[0])
	rows := len(matrix1)

	subMatrix := CreateZeroMatrix[T](rows, cols)
	fmt.Println(subMatrix)
	
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			subMatrix[row][col] = matrix1[row][col] - matrix2[row][col]
		}
	}
	return subMatrix, nil
}
func DotProduct[T Matrix](matrix1, matrix2 [][]T) ([][]T, error) {
	row1 := len(matrix1)
	col1 := len(matrix1[0])

	row2 := len(matrix2)
	col2 := len(matrix2[0])

	if col1 != row2 {
		return nil, fmt.Errorf("matrix dimensions cannot be multiplied! Wrong dimension: %dx%d & %dx%d", row1, col1, row2, col2)
	}

	prodMatrix := CreateZeroMatrix[T](row1, col2)
	
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
func TransposeMatrix[T Matrix](matrix [][]T) [][]T {

	rows := len(matrix[0])
	cols := len(matrix)

	tMatrix := CreateZeroMatrix[T](rows, cols)
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			tMatrix[row][col] = matrix[col][row]
		}
	}
	
	
	return tMatrix
}
func CreateZeroMatrix[T Matrix](rows, cols int) [][]T {
	matrix := make([][]T, rows)
	for i := range matrix {
		matrix[i] = make([]T, cols)
	}
	return matrix
}
func PrintMatrix[T Matrix](matrixName string, matrix [][]T) {
	fmt.Println()
	fmt.Println(matrixName, ": ", len(matrix), "x", len(matrix[0]))
	for _, row := range matrix {
		fmt.Println(row);
	}
}