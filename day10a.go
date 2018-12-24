package main

import (
  "bufio"
  "fmt"
  "os"
)

type Star struct {
  X int
  Y int
  VX int
  VY int
}

func extents(stars []*Star) (int, int) {
  xmax := 0
  ymax := 0
  for i := 0; i < len(stars); i++ {
    s := stars[i]
    if s.X > xmax {
      xmax = s.X
    }
    if s.Y > ymax {
      ymax = s.Y
    }
  }
  return xmax, ymax
}

func step(stars []*Star) {
  for i := 0; i < len(stars); i++ {
    s := stars[i]
    s.X += s.VX
    s.Y += s.VY
  }
}

func rewind(stars []*Star) {
  for i := 0; i < len(stars); i++ {
    s := stars[i]
    s.X -= s.VX
    s.Y -= s.VY
  }
}

func display(stars []*Star) {
  xmax, ymax := extents(stars)
  fmt.Println(xmax, ymax)
  for y := 0; y <= ymax; y++ {
    for x := 0; x <= xmax; x++ {
      starHere := false
      for _, s := range stars {
        if s.X == x && s.Y == y {
          starHere = true
        }
      }
      if starHere {
        fmt.Print("#")
      } else {
        fmt.Print(".")
      }
    }
    fmt.Print("\n")
  }
  fmt.Print("\n")
  return
}

func readInput() []*Star {
  result := make([]*Star, 0)
  reader := bufio.NewReader(os.Stdin)
  bytes, _, err := reader.ReadLine()
  for err == nil {
    in := string(bytes)
    s := Star{}
    fmt.Sscanf(in, "position=<%d, %d> velocity=<%d, %d>", &s.X, &s.Y, &s.VX, &s.VY)
    result = append(result, &s)
    bytes, _, err = reader.ReadLine()
  }
  return result
}

func iterate(stars []*Star) int {
  bestsum := 100000000
  elapsed := 0
  for i := 0; i < 100000; i++ {
    x, y := extents(stars)
    if x + y < bestsum {
      bestsum = x + y
    } else {
      rewind(stars)
      display(stars)
      elapsed--
      return elapsed
    }
    step(stars)
    elapsed++
  }
  return -1
}

func main() {
  stars := readInput()
  elapsed := iterate(stars)
  fmt.Println(elapsed)
}
