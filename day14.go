package main

import (
  "fmt"
  "os"
  "container/list"
  "strconv"
)

//func printBoard(board list) {
//    b := board.Front()
//    for i := 0; i < board.Len(); i++ {
//     fmt.Print(b.Value.(int), " ")
//     b = b.Next()
//    }
//    fmt.Println("")
//}

func process(rounds int) []int {
  board := list.New()
  elf1 := board.PushBack(3)
  elf2 := board.PushBack(7)
  for {
    score1 := elf1.Value.(int)
    score2 := elf2.Value.(int)
    newRecipe := score1 + score2
    if newRecipe >= 10 {
      board.PushBack(1)
    }
    board.PushBack(newRecipe % 10)
    for e1 := 0; e1 < score1+1; e1++ {
      elf1 = elf1.Next()
      if elf1 == nil {
        elf1 = board.Front()
      }
    }
    for e2 := 0; e2 < score2+1; e2++ {
      elf2 = elf2.Next()
      if elf2 == nil {
        elf2 = board.Front()
      }
    }
    //printBoard(board)

    if board.Len() > rounds + 10 {
      c := board.Front()
      for i := 0; i < rounds; i++ {
        c = c.Next()
      }
      result := make([]int, 0)
      for i := 0; i < 10; i++ {
        result = append(result, c.Value.(int))
        c = c.Next()
      }
      return result
    }
  }
}

func main() {
  rounds, _ := strconv.Atoi(os.Args[1])
  result := process(rounds)
  fmt.Println(result)
}

