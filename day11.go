package main

import (
  "fmt"
  "os"
  "strconv"
)

func powerLevel(x int, y int, serial int) int {
  rackID := x + 10
  result := rackID * y
  result += serial
  result *= rackID
  result = result % 1000 / 100
  result -= 5
  return result
}

func bestCell(grid [][]int) (int, int) {
  bestX := 0
  bestY := 0
  bestPower := -1000000
  for y := 0; y < 298; y++ {
    for x := 0; x < 298; x++ {
      p := grid[y][x] + grid[y][x+1] + grid[y][x+2] +
        grid[y+1][x] + grid[y+1][x+1] + grid[y+1][x+2] +
        grid[y+2][x] + grid[y+2][x+1] + grid[y+2][x+2]
      if p > bestPower {
        bestPower = p
        bestX = x
        bestY = y
      }
    }
  }
  return bestX+1, bestY+1
}

func maxPower(serial int) (int, int) {
  grid := make([][]int, 300)
  for i := 0; i < 300; i++ {
    grid[i] = make([]int, 300)
  }
  for y := 0; y < 300; y++ {
    for x := 0; x < 300; x++ {
      grid[y][x] = powerLevel(x+1, y+1, serial)
    }
  }
  bestX, bestY := bestCell(grid)
  return bestX, bestY
}

func main() {
  serial, _ := strconv.Atoi(os.Args[1])
  x, y := maxPower(serial)
  fmt.Println(x, ",", y)
  // test cases
  //fmt.Println(powerLevel(3, 5, 8))
  //fmt.Println(powerLevel(122, 79, 57))
  //fmt.Println(powerLevel(217, 196, 39))
  //fmt.Println(powerLevel(101, 153, 71))
}
