package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {

	f, _ := os.Open("input.txt")
	/* sum := partOne(f)
	fmt.Println("Part1: ", sum) */

	// Reset the file pointer
	f.Seek(0, 0)

	p2 := partTwo(f)
	fmt.Println("Part2: ", p2)

}

func partOne(file io.Reader) int {
	// Json decoder from the file. Will allow us to go token by token. Go :)
	data := json.NewDecoder(file)
	sum := 0.0

	for {
		// While there's data to read and tokenize from
		if value, err := data.Token(); err == nil {
			if reflect.TypeOf(value).Kind() == reflect.Float64 {
				// fmt.Println("Value: ", value)
				sum += value.(float64)
			}
		} else {
			// No more data
			break
		}
	}

	return int(sum)
}

func partTwo(file io.Reader) int {

	scanner := bufio.NewScanner(file)
	bytes := []byte{}

	for scanner.Scan() {
		// The file is one big line of bytes
		bytes = scanner.Bytes()
	}

	data := map[string]any{}
	json.Unmarshal(bytes, &data)

	return sumItem(data)
}

func sumItem(element any) int {
	switch reflect.TypeOf(element).Kind() {
	case reflect.Map:
		return sumMap(element.(map[string]any))
	case reflect.Slice:
		return sumArray(element.([]any))
	case reflect.Float64:
		return int(element.(float64))
	default:
		return 0
	}
}

// Sum a Map. discard maps with red as a value
func sumMap(dict map[string]any) int {
	sum := 0

	for _, value := range dict {
		// If the value of the keys is red
		// Also check if the type assertion is successful
		if s, ok := value.(string); ok && s == "red" {
			return 0
		}
		sum += sumItem(value)
	}
	return sum
}

// An array may contain numbers only, strings only, and mix match
func sumArray(numbers []any) int {
	sum := 0
	for _, n := range numbers {
		sum += sumItem(n)
	}

	return sum
}
