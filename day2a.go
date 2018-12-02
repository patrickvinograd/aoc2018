package main

import (
	"fmt"
)

func readInput() []string {
	result := make([]string, 0)
	var in string
	lines, _ := fmt.Scanln(&in)
	for lines > 0 {
		result = append(result, in)
		lines, _ = fmt.Scanln(&in)
	}
	return result
}

func delta(s1 string, s2 string) int {
	delta := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			delta++
		}
	}
	return delta
}

func common(s1 string, s2 string) string {
	result := make([]byte, 0)
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			result = append(result, s1[i])
		}
	}
	return string(result)
}

func match(values []string) string {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if delta(values[i], values[j]) == 1 {
				return common(values[i], values[j])
			}
		}
	}
	return "not found"
}

func main() {
	values := readInput()
	answer := match(values)
	fmt.Println(answer)
}
