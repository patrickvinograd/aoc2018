package main

import (
	"fmt"
	"strings"
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

func hasTwos(chars string) (bool, bool) {
	twos, threes := false, false
	for len(chars) > 0 {
		orig := len(chars)
		chars = strings.Replace(chars, chars[0:1], "", -1)
		new := len(chars)
		if orig-new == 2 {
			twos = true
		} else if orig-new == 3 {
			threes = true
		}
	}
	return twos, threes
}

func checksum(values []string) int {
	twos, threes := 0, 0
	for i := 0; i < len(values); i++ {
		val := values[i]
		h2, h3 := hasTwos(val)
		if h2 == true {
			twos++
		}
		if h3 == true {
			threes++
		}
	}
	return twos * threes
}

func main() {
	values := readInput()
	answer := checksum(values)
	fmt.Println(answer)
}
