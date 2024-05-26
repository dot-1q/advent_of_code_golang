package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	rules := createRules()
	part := 2

	m := [][]rune{
		{'.', '#', '.'},
		{'.', '.', '#'},
		{'#', '#', '#'}}
	if part == 1 {
		m = enhance(m, rules, 5)
		fmt.Println("Part 1 : ", howManyOn(m))
	} else {
		m = enhance(m, rules, 18)
		fmt.Println("Part 2 : ", howManyOn(m))
	}
}

func enhance(matrix [][]rune, rules map[string]string, times int) [][]rune {
	m := matrix
	for range times {
		result := [][][]rune{}
		if len(m)%2 == 0 {
			// Array of 2D arrays
			twod := extract2x2Matrices(m, len(m))
			for _, square := range twod {
				// Resulting matrices from the transformation.
				result = append(result, applyTransformation(square, rules))
			}
			// The 2by2 arrays get transformed into multiple 3x3
			m = merge3x3Matrices(result)
		} else {
			treed := extract3x3Matrices(m, len(m))
			for _, square := range treed {
				result = append(result, applyTransformation(square, rules))
			}
			// The 3by3 arrays get transformed into multiple 4x4
			m = merge4x4Matrices(result)
		}
	}
	return m
}

// Function to merge 4x4 matrices into a larger matrix
func merge4x4Matrices(matrices [][][]rune) [][]rune {
	numMatrices := len(matrices)
	dim := int(math.Sqrt(float64(numMatrices))) * 4
	largeMatrix := make([][]rune, dim)
	for i := range largeMatrix {
		largeMatrix[i] = make([]rune, dim)
	}

	// Fill the larger matrix with the 4x4 matrices
	for idx, matrix := range matrices {
		rowOffset := (idx / (dim / 4)) * 4
		colOffset := (idx % (dim / 4)) * 4
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				largeMatrix[rowOffset+i][colOffset+j] = matrix[i][j]
			}
		}
	}
	return largeMatrix
}

func merge3x3Matrices(matrices [][][]rune) [][]rune {
	// Calculate the dimension of the larger matrix
	numMatrices := len(matrices)
	dim := int(math.Sqrt(float64(numMatrices))) * 3
	largeMatrix := make([][]rune, dim)
	for i := range largeMatrix {
		largeMatrix[i] = make([]rune, dim)
	}

	// Fill the larger matrix with the 3x3 matrices
	for idx, matrix := range matrices {
		rowOffset := (idx / (dim / 3)) * 3
		colOffset := (idx % (dim / 3)) * 3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				largeMatrix[rowOffset+i][colOffset+j] = matrix[i][j]
			}
		}
	}
	return largeMatrix
}

func applyTransformation(matrix [][]rune, rules map[string]string) [][]rune {
	// Get the string representation of this matrix
	str := MatrixToStr(matrix)
	transformation := rules[str]
	return StrToMatrix(transformation)
}

func extract2x2Matrices(matrix [][]rune, N int) [][][]rune {
	var result [][][]rune

	for i := 0; i < N; i += 2 {
		for j := 0; j < N; j += 2 {
			// Create a new 2x2 matrix
			newMatrix := [][]rune{
				{matrix[i][j], matrix[i][j+1]},
				{matrix[i+1][j], matrix[i+1][j+1]},
			}
			result = append(result, newMatrix)
		}
	}
	return result
}

func extract3x3Matrices(matrix [][]rune, N int) [][][]rune {
	var result [][][]rune

	for i := 0; i < N; i += 3 {
		for j := 0; j < N; j += 3 {
			// Create a new 3x3 matrix
			newMatrix := [][]rune{
				{matrix[i][j], matrix[i][j+1], matrix[i][j+2]},
				{matrix[i+1][j], matrix[i+1][j+1], matrix[i+1][j+2]},
				{matrix[i+2][j], matrix[i+2][j+1], matrix[i+2][j+2]},
			}
			result = append(result, newMatrix)
		}
	}

	return result
}

func createRules() map[string]string {
	rules := map[string]string{}

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=>")
		// Don't forget to trim space
		equivRules := createRotations(strings.TrimSpace(line[0]))
		// For every equivalent rotation, associate it with the transformation.
		for _, rule := range equivRules {
			rules[rule] = strings.TrimSpace(line[1])
		}
	}
	return rules
}

func createRotations(rule string) []string {
	// Add the original rule
	rotations := []string{rule}
	// Create the matrix from the rule
	matrix := StrToMatrix(rule)

	// Apply the all transformations, and return to string format and append.
	rotate90 := rotateClockwise(matrix)
	rotations = append(rotations, MatrixToStr(rotate90))
	rotate180 := rotateClockwise(rotate90)
	rotations = append(rotations, MatrixToStr(rotate180))
	rotate270 := rotateClockwise(rotate180)
	rotations = append(rotations, MatrixToStr(rotate270))

	// Flip and apply the rotations on the flipped version.
	flipH := flipHorizontal(matrix)
	rotations = append(rotations, MatrixToStr(flipH))
	flipRot90 := rotateClockwise(flipH)
	rotations = append(rotations, MatrixToStr(flipRot90))
	flipRot180 := rotateClockwise(flipRot90)
	rotations = append(rotations, MatrixToStr(flipRot180))
	flipRot270 := rotateClockwise(flipRot180)
	rotations = append(rotations, MatrixToStr(flipRot270))

	return rotations
}

func StrToMatrix(rule string) [][]rune {
	matrix := [][]rune{}

	rows := strings.Split(rule, "/")
	for _, row := range rows {
		r := []rune{}
		for _, char := range row {
			r = append(r, char)
		}
		matrix = append(matrix, r)
	}
	return matrix
}

func MatrixToStr(matrix [][]rune) string {
	s := strings.Builder{}

	for i, row := range matrix {
		for _, char := range row {
			s.WriteRune(char)
		}
		if i != len(matrix)-1 {
			s.WriteRune('/')
		}
	}
	return s.String()
}

// Rotate but create a new matrix in the process.
func rotateClockwise(matrix [][]rune) [][]rune {
	// Create new matrix
	m := make([][]rune, len(matrix))
	for i := range len(matrix) {
		m[i] = make([]rune, len(matrix[i]))
		copy(m[i], matrix[i])
	}

	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = m[j], m[i]
	}

	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	return m
}

func flipHorizontal(matrix [][]rune) [][]rune {
	// Create new matrix
	m := make([][]rune, len(matrix))
	for i := range len(matrix) {
		m[i] = make([]rune, len(matrix[i]))
		copy(m[i], matrix[i])
	}
	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = matrix[j], matrix[i]
	}
	return m
}

func howManyOn(m [][]rune) int {
	s := 0
	for _, row := range m {
		for _, char := range row {
			if char == '#' {
				s++
			}
		}
	}
	return s
}
