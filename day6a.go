package main

import (
	"bufio"
	"fmt"
	"os"
)

type Place struct {
	Id int
	X  int
	Y  int
}

type Coord struct {
	Prox int
	X    int
	Y    int
}

func readInput() []Place {
	result := make([]Place, 0)
	reader := bufio.NewReader(os.Stdin)
	bytes, _, err := reader.ReadLine()
	id := 0
	for err == nil {
		in := string(bytes)
		p := Place{Id: id}
		id++
		fmt.Sscanf(in, "%d, %d", &p.X, &p.Y)
		result = append(result, p)
		bytes, _, err = reader.ReadLine()
	}
	return result
}

func buildMap(places []Place) [][]Coord {
	xmax, ymax := 0, 0
	for _, i := range places {
		if i.X > xmax {
			xmax = i.X
		}
		if i.Y > ymax {
			ymax = i.Y
		}
	}
	m := make([][]Coord, ymax+1)
	for i := 0; i <= ymax; i++ {
		m[i] = make([]Coord, xmax+1)
		for j := 0; j <= xmax; j++ {
			m[i][j] = Coord{X: j, Y: i}
		}
	}
	return m
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func totalDist(x int, y int, places []Place) int {
	total := 0
	for _, p := range places {
		d := Abs(p.X-x) + Abs(p.Y-y)
		total = total + d
	}
	return total
}

func distances(m [][]Coord, places []Place) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			id := totalDist(x, y, places)
			m[y][x].Prox = id
		}
	}
}

func safeCount(m [][]Coord) int {
	result := 0
	for _, row := range m {
		for _, place := range row {
			if place.Prox < 10000 {
				result++
			}
		}
	}
	return result
}

func main() {
	places := readInput()
	fmt.Println(places)
	m := buildMap(places)
	distances(m, places)
	fmt.Println(safeCount(m))
}
