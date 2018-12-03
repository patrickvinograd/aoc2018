package main

import (
	"fmt"
)

type Claim struct {
	Id     int
	X      int
	Y      int
	Width  int
	Height int
}

func readInput() []Claim {
	result := make([]Claim, 0)
	c := Claim{}
	lines, _ := fmt.Scanf("#%d @ %d,%d: %dx%d\n", &c.Id, &c.X, &c.Y, &c.Width, &c.Height)
	for lines > 0 {
		result = append(result, c)
		c = Claim{}
		lines, _ = fmt.Scanf("#%d @ %d,%d: %dx%d\n", &c.Id, &c.X, &c.Y, &c.Width, &c.Height)
	}
	return result
}

func extent(claims []Claim) (int, int) {
	xmax, ymax := 0, 0
	for i := 0; i < len(claims); i++ {
		c := claims[i]
		x := c.X + c.Width
		y := c.Y + c.Height
		if x > xmax {
			xmax = x
		}
		if y > ymax {
			ymax = y
		}
	}
	return xmax + 1, ymax + 1
}

func overlap(claims []Claim) int {
	x, y := extent(claims)
	fabric := make([][]int, x)
	for i := 0; i < x; i++ {
		fabric[i] = make([]int, y)
	}

	for i := 0; i < len(claims); i++ {
		c := claims[i]
		for j := c.X; j < c.X+c.Width; j++ {
			for k := c.Y; k < c.Y+c.Height; k++ {
				fabric[j][k] = fabric[j][k] + 1
			}
		}
	}
	result := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if fabric[i][j] > 1 {
				result++
			}
		}
	}
	return result

}

func main() {
	values := readInput()
	result := overlap(values)
	fmt.Println(result)
}
