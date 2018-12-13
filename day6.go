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

func minDist(x int, y int, places []Place) int {
	min := 1000000
	dup := false
	best := -1
	for _, p := range places {
		d := Abs(p.X-x) + Abs(p.Y-y)
		if d < min {
			min = d
			best = p.Id
			dup = false
		} else if d == min {
			dup = true
		}
	}
	if dup {
		return -1
	}
	return best
}

func distances(m [][]Coord, places []Place) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			id := minDist(x, y, places)
			m[y][x].Prox = id
		}
	}
}

func biggestArea(m [][]Coord) int {
	counts := make(map[int]int)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			id := m[i][j].Prox
			counts[id] = counts[id] + 1
		}
	}
	fmt.Println(counts)
	exclude := make(map[int]bool)
	for _, i := range m[0] {
		exclude[i.Prox] = true
	}
	for _, i := range m[len(m)-1] {
		exclude[i.Prox] = true
	}
	for i := 0; i < len(m); i++ {
		exclude[m[i][0].Prox] = true
		exclude[m[i][len(m[i])-1].Prox] = true
	}
	fmt.Println(exclude)
	for k := range exclude {
		delete(counts, k)
	}
	max := 0
	for k := range counts {
		if counts[k] > max {
			max = counts[k]
		}
	}
	return max
}

func main() {
	places := readInput()
	fmt.Println(places)
	m := buildMap(places)
	distances(m, places)
	for _, row := range m {
		fmt.Println(row)
	}
	biggest := biggestArea(m)
	fmt.Println(biggest)
	//fmt.Println(m)
}
