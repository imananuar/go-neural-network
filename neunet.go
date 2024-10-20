package main

// import "fmt"

//Forward Propagation
func ActFunction[T Matrix](weightage, input [][]T, bias T) ([][]T, error) {

	z, err := DotProduct(weightage, input)
	if err != nil {
		return nil, err
	}

	rowz := len(z)
	colz := len(z[0])
	biasMatrix := CreateZeroMatrix[T](rowz, colz)
	
	for row := 0; row < len(biasMatrix); row++ {
		for col := 0; col < len(biasMatrix[0]); col++ {
			biasMatrix[row][col] = bias;
		}
	}

	neuron, err := AddMatrix[T](z, biasMatrix);
	if err != nil {
		return nil, err
	}
	return neuron, nil
}