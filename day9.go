package main

import (
	"bufio"
	"fmt"
	"os"
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

func delete(vPtr *[]int, i int) int {
  //fmt.Println("delete", i)
  v := *vPtr
  result := v[i]
  copy(v[i:], v[i+1:])
  v[len(v)-1] = 0 // or the zero value of T
  *vPtr = v[:len(v)-1]
  return result
}

func insert(vPtr *[]int, i int, x int) {
  v := *vPtr
  v = append(v, 0)
  copy(v[i+1:], v[i:])
  v[i] = x
  *vPtr = v
}

func play(playerCount int, highMarble int) map[int]int {
  players := make([]int, playerCount)
  for i := 0; i < playerCount; i++ {
    players[i] = i+1
  }
  scores := make(map[int]int)

  turn := 0
  current := 0
  circle := make([]int, 1)
  circle[current] = 0

  for marble := 1; marble <= highMarble; marble++ {
    if marble % 23 == 0 {
      id := players[turn]
      removeAt := (current - 7 + len(circle)) % len(circle)
      //fmt.Println("deleting", current, removeAt, len(circle))
      removed := delete(&circle, removeAt)
      if removeAt == len(circle) {
        current = 0
      } else {
        current = removeAt
      }
      scores[id] = scores[id] + marble + removed
    } else {
      insertAt := (current + 2) % len(circle)
      if insertAt == 0 {
        circle = append(circle, marble)
        current = len(circle) - 1
      } else {
        insert(&circle, insertAt, marble)
        current = insertAt
      }
    }
    //fmt.Println(marble, turn, current, circle[current])
    //fmt.Println(circle)
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
  fmt.Println(players, "players", highMarble, "marble")
  scores := play(players, highMarble)
  fmt.Println(highScore(scores))
}
