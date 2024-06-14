package main

import "fmt"

func transposeMatrix(matrix *[][]int) {
	rows := len(*matrix)
	cols := len((*matrix)[0])

	transposed := make([][]int, cols)

	for i := range transposed {
		transposed[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = (*matrix)[i][j]
		}
	}

	*matrix = transposed
}

func main() {

	var matrix [][]int = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("Original matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
	transposeMatrix(&matrix)
	fmt.Println("Transpose matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
