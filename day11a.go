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

func powerTotal(grid [][]int, xp int, yp int, size int) int {
  result := 0
  for y := yp; y < yp + size; y++ {
    for x := xp; x < xp + size; x++ {
      result += grid[y][x]
    }
  }
  return result
}

func bestCell(grid [][]int) (int, int, int) {
  bestX := 0
  bestY := 0
  bestSize := 0
  bestPower := -1000000
  for size := 1; size <= 300; size++ {
    ymax := 301 - size
    xmax := 301 - size
    for y := 0; y < ymax; y++ {
      for x := 0; x < xmax; x++ {
        p := powerTotal(grid, x, y, size)
        if p > bestPower {
          bestPower = p
          bestX = x
          bestY = y
          bestSize = size
        }
      }
    }
  }
  fmt.Println(bestPower)
  return bestX+1, bestY+1, bestSize
}

func maxPower(serial int) (int, int, int) {
  grid := make([][]int, 300)
  for i := 0; i < 300; i++ {
    grid[i] = make([]int, 300)
  }
  for y := 0; y < 300; y++ {
    for x := 0; x < 300; x++ {
      grid[y][x] = powerLevel(x+1, y+1, serial)
    }
  }
  bestX, bestY, bestSize := bestCell(grid)
  return bestX, bestY, bestSize
}

func main() {
  serial, _ := strconv.Atoi(os.Args[1])
  x, y, size := maxPower(serial)
  fmt.Println(x, ",", y, ",", size)
  // test cases
  //fmt.Println(powerLevel(3, 5, 8))
  //fmt.Println(powerLevel(122, 79, 57))
  //fmt.Println(powerLevel(217, 196, 39))
  //fmt.Println(powerLevel(101, 153, 71))
}
