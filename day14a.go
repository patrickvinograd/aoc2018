package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

//func printBoard(board list) {
//    b := board.Front()
//    for i := 0; i < board.Len(); i++ {
//     fmt.Print(b.Value.(int), " ")
//     b = b.Next()
//    }
//    fmt.Println("")
//}

func indexOf(input []int, substr []int) int {
  result := -1
  start := len(input) - len(substr) -1;
  if start < 0 {
    start = 0
  } 
  for i:= start; i < len(input) - len(substr); i++ {
    //fmt.Println(len(input), len(substr))
    match := true
    for j := 0; j < len(substr); j++ {
      if input[i+j] != substr[j] {
        match = false
      }
    }
    if match {
      return i
    }
  }
  return result
}

func process(target []int) int {
  elf1 := 0
  elf2 := 1
  board := make([]int, 0)
  board = append(board, 3)
  board = append(board, 7)
  for {
    score1 := board[elf1]
    score2 := board[elf2]
    newRecipe := score1 + score2
    if newRecipe >= 10 {
      board = append(board, 1)
    }
    board = append(board, newRecipe % 10)
    for e1 := 0; e1 < score1+1; e1++ {
      elf1 = elf1 + 1
      if elf1 >= len(board) {
        elf1 = 0
      }
    }
    for e2 := 0; e2 < score2+1; e2++ {
      elf2 = elf2 + 1
      if elf2 >= len(board) {
        elf2 = 0
      }
    }
    //fmt.Println(elf1, elf2)
    //fmt.Println(board)

    if i := indexOf(board, target); i >= 0 {
      return i
    }
  }
}

func split(targetStr string) []int {
  digits := strings.Split(targetStr, "")
  result := make([]int, len(digits))
  for i := 0; i < len(digits); i++ {
    result[i], _ = strconv.Atoi(digits[i])
  }
  return result
}

func main() {
  target := split(os.Args[1])
  fmt.Println(target)
  result := process(target)
  fmt.Println(result)
}

