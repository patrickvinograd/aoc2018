package main

import (
	"fmt"
)

type Claim struct {
	Id      int
	X       int
	Y       int
	Width   int
	Height  int
	Overlap bool
}

type Space struct {
	X      int
	Y      int
	Count  int
	Claims []*Claim
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
	fabric := make([][]*Space, x)
	for i := 0; i < x; i++ {
		fabric[i] = make([]*Space, y)
		for j := 0; j < y; j++ {
			fabric[i][j] = &Space{i, j, 0, make([]*Claim, 0)}
		}
	}

	for i := 0; i < len(claims); i++ {
		c := &claims[i]
		for j := c.X; j < c.X+c.Width; j++ {
			for k := c.Y; k < c.Y+c.Height; k++ {
				s := fabric[j][k]
				s.Count = s.Count + 1
				s.Claims = append(s.Claims, c)
				if s.Count > 1 {
					for x := 0; x < len(s.Claims); x++ {
						(*s.Claims[x]).Overlap = true
					}
				}
			}
		}
	}
	for i := 0; i < len(claims); i++ {
		c := &claims[i]
		if c.Overlap == false {
			fmt.Println(*c)
			return c.Id
		}
	}
	return -1
}

func main() {
	values := readInput()
	result := overlap(values)
	fmt.Println(result)
}
