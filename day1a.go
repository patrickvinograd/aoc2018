package main

import (
	"fmt"
	"strconv"
)

func readInput() []int {
	result := make([]int, 0)
	var in string
	lines, _ := fmt.Scanln(&in)
	for lines > 0 {
		freq, _ := strconv.Atoi(in)
		result = append(result, freq)
		lines, _ = fmt.Scanln(&in)
	}
	return result
}

func findRepeat(values []int) int {
	seen := make(map[int]bool)
	counter, i := 0, 0
	for {
		counter, i = counter+values[i], i+1
		if seen[counter] == true {
			return counter
		} else {
			seen[counter] = true
		}
		if i >= len(values) {
			i = 0
		}
	}
}

func main() {
	values := readInput()
	//fmt.Println(values)
	answer := findRepeat(values)
	fmt.Println(answer)
}
