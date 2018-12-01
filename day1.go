package main

import (
  "fmt"
  "strconv"
)

func main() {
  var result int = 0
  var in string
  lines, _ := fmt.Scanln(&in)
  for lines > 0 {
    freq, _ := strconv.Atoi(in)
    result += freq
    lines, _ = fmt.Scanln(&in)
  }
  fmt.Println(result)
}
