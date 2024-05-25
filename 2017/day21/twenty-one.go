package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// m1 := [][]rune{{'.', '#', '.'}, {'.', '.', '#'}, {'#', '#', '#'}}
	rules := createRules()

	m2 := [][]rune{
		{'.', '#', '.'},
		{'.', '.', '#'},
		{'#', '#', '#'}}
	for range 4 {
		for i := range m2 {
			fmt.Printf("%c\n", m2[i])
		}
		m2 = enhance(m2, rules)
		fmt.Println("----")
		for i := range m2 {
			fmt.Printf("%c\n", m2[i])
		}
		fmt.Println("Done")
	}

	// for k := range rules {
	// 	fmt.Printf("%s: => %s\n", k, rules[k])
	// }
}

func enhance(matrix [][]rune, rules map[string]string) [][]rune {
	if len(matrix)%2 == 0 {
		fmt.Println("Here")
		// Array of 2D arrays
		twod := twobytwo(matrix)
		result := [][][]rune{}
		for _, square := range twod {
			// Resulting matrices from the transformation.
			result = append(result, applyTransformation(square, rules))
		}
		return concatMatrixTwo(result)
	} else {
		treed := threebythree(matrix)
		result := [][][]rune{}
		for _, square := range treed {
			result = append(result, applyTransformation(square, rules))
		}
		return concatMatrixThree(result)
	}
}

func applyTransformation(matrix [][]rune, rules map[string]string) [][]rune {
	// Get the string representation of this matrix
	str := MatrixToStr(matrix)
	transformation := rules[str]
	return createMatrix(transformation)
}

func concatMatrixTwo(matrices [][][]rune) [][]rune {
	m := make([][]rune, len(matrices[0])*2)
	for _, g := range matrices {
		for _, row := range g {
			fmt.Printf("%c\n", row)
		}
		fmt.Println("--------")
	}

	for i := 0; i < len(matrices); i = i + 3 {
		c := 0
		r := matrices[c][:][0]
		x := matrices[c+1][:][0]
		m[i] = append(m[i], r...)
		m[i] = append(m[i], x...)

		r2 := matrices[c][:][1]
		x2 := matrices[c+1][:][1]
		m[i+1] = append(m[i+1], r2...)
		m[i+1] = append(m[i+1], x2...)

		r3 := matrices[c][:][2]
		x3 := matrices[c+1][:][2]
		m[i+2] = append(m[i+2], r3...)
		m[i+2] = append(m[i+2], x3...)

	}
	return m
}

func concatMatrixThree(matrices [][][]rune) [][]rune {
	m := make([][]rune, len(matrices[0]))
	if len(matrices) == 1 {
		return matrices[0]
	}

	for i := 0; i < len(m); i = i + 3 {
		r := matrices[i][:][0]
		x := matrices[i+1][:][0]
		y := matrices[i+2][:][0]
		m[i] = append(m[i], r...)
		m[i] = append(m[i], x...)
		m[i] = append(m[i], y...)

		r2 := matrices[i][:][1]
		x2 := matrices[i+1][:][1]
		y2 := matrices[i+2][:][1]
		m[i+1] = append(m[i+1], r2...)
		m[i+1] = append(m[i+1], x2...)
		m[i+1] = append(m[i+1], y2...)

		r3 := matrices[i][:][2]
		x3 := matrices[i+1][:][2]
		y3 := matrices[i+2][:][2]
		m[i+2] = append(m[i+2], r3...)
		m[i+2] = append(m[i+2], x3...)
		m[i+2] = append(m[i+2], y3...)
	}
	return m
}

func twobytwo(matrix [][]rune) [][][]rune {
	// This is such insane slicing, that if you want to understand just print along the way
	// It just breaks down 4x4,6x6,8x8,ect... into 2x2 squares.
	results := [][][]rune{}
	for i := 0; i < len(matrix); i = i + 2 {
		for cell := 0; cell < len(matrix[i]); cell = cell + 2 {
			g := matrix[i : i+2]
			n := make([][]rune, len(g))
			// This is needed, as slicing [][x:x+2] doesnt do what it should, i.e, slicing the inner array.
			for x, inner := range g {
				n[x] = inner[cell : cell+2]
			}
			results = append(results, n)
		}
	}
	return results
}

func threebythree(matrix [][]rune) [][][]rune {
	results := [][][]rune{}
	for i := 0; i < len(matrix); i = i + 3 {
		for cell := 0; cell < len(matrix[i]); cell = cell + 3 {
			g := matrix[i : i+3]
			n := make([][]rune, len(g))
			for x, inner := range g {
				n[x] = inner[cell : cell+3]
			}
			results = append(results, n)
		}
	}
	return results
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
	matrix := createMatrix(rule)

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

func createMatrix(rule string) [][]rune {
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

func flipVertical(matrix [][]rune) [][]rune {
	// Create new matrix
	m := make([][]rune, len(matrix))
	for i := range len(matrix) {
		m[i] = make([]rune, len(matrix[i]))
		copy(m[i], matrix[i])
	}
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			m[i][j], m[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	m = rotateClockwise(m)
	return m
}
