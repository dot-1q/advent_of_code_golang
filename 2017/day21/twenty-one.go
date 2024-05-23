package main

import "fmt"

func main() {
	m1 := [][]rune{{'.', '#', '.'}, {'.', '.', '#'}, {'#', '#', '#'}}
	for i := range m1 {
		fmt.Printf("%c\n", m1[i])
	}
	fmt.Println("--------------------")
	// rotateClockwise(m1)
	// rotateAntiClockwise(m1)
	flipHorizontal(m1)
	// flipVertical(m1)
	// flipVertical(m1)

	for i := range m1 {
		fmt.Printf("%c\n", m1[i])
	}
}

func rotateClockwise(matrix [][]rune) [][]rune {
	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func rotateAntiClockwise(matrix [][]rune) [][]rune {
	// reverse the matrix
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// transpose it
	i := 0
	j := 0
	column := 0
	for column < len(matrix) {
		i = 0
		j = len(matrix) - 1
		for i < j {
			matrix[i][column], matrix[j][column] = matrix[j][column], matrix[i][column]
			i = i + 1
			j = j - 1
		}
		column = column + 1
	}
	return matrix
}

func flipHorizontal(matrix [][]rune) [][]rune {
	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
	return matrix
}

func flipVertical(matrix [][]rune) [][]rune {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	rotateClockwise(matrix)
	return matrix
}
