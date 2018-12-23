package main

import (
	"bufio"
	"fmt"
	"os"
  "container/list"
)

func readInput() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
  var players int
  var highMarble int
  in := string(bytes)
  fmt.Sscanf(in, "%d players; last marble is worth %d points", &players, &highMarble)
  return players, highMarble
}

func play(playerCount int, highMarble int) map[int]int {
  players := make([]int, playerCount)
  for i := 0; i < playerCount; i++ {
    players[i] = i+1
  }
  scores := make(map[int]int)

  turn := 0
  circle := list.New()
  current := circle.PushBack(0)

  for marble := 1; marble <= highMarble; marble++ {
    if marble % 1000 == 0 {
      fmt.Println(marble)
    }
    if marble % 23 == 0 {
      id := players[turn]
      for x:= 0; x < 7; x++ {
        current = current.Prev()
        if current == nil {
          current = circle.Back()
        }
      }
      toRemove := current
      current = current.Next()
      if current == nil {
        current = circle.Front()
      }
      removed := circle.Remove(toRemove).(int)
      scores[id] = scores[id] + marble + removed
    } else {
      current = current.Next()
      if current == nil {
        current = circle.Front()
      }
      current = circle.InsertAfter(marble, current)
    }
    turn++
    if turn >= len(players) {
      turn = 0
    }
  }
  return scores
}

func highScore(scores map[int]int) int {
  result := 0
  for _, v := range scores {
    if v > result {
      result = v
    }
  }
  return result
}

func main() {
	players, highMarble := readInput()
  fmt.Println(players, "players", highMarble*100, "marble")
  scores := play(players, highMarble*100)
  fmt.Println(highScore(scores))
}
