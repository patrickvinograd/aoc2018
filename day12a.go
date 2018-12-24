package main

import (
  "fmt"
  "os"
  "bufio"
)

func readInput() (string, map[string]byte) {
  reader := bufio.NewReader(os.Stdin)
  bytes, _, err := reader.ReadLine()
  in := string(bytes)
  istate := ""
  fmt.Sscanf(in, "initial state: %s", &istate)
  transitions := make(map[string]byte)
  _, _, _ = reader.ReadLine()
  bytes, _, err = reader.ReadLine()
  for err == nil {
    in := string(bytes)
    from := ""
    to := ""
    fmt.Sscanf(in, "%5s => %1s", &from, &to)
    if to == "#" {
      transitions[from] = '#'
    } else {
      transitions[from] = '.'
    }
    transitions[from] = to[0]
    bytes, _, err = reader.ReadLine()
  }
  return istate, transitions
}

func initialize(istate string) []byte {
  result := []byte("........................................")
  sb := []byte(istate)
  for i := 0; i < len(sb); i++ {
    result = append(result, sb[i])
  }
  for i := 0; i < 500; i++ {
    result = append(result, '.')
  }
  return result
}

func generation(prev []byte, rules map[string]byte) []byte {
  next := make([]byte, len(prev))
  for i := 2; i < len(prev)-2; i++ {
    rule := string(prev[i-2:i+3])
    out := rules[rule]
    if out == 0 {
      out = '.'
    }
    next[i] = out
  }
  next[0] = prev[0]
  next[1] = prev[1]
  next[len(prev)-1] = prev[len(prev)-1]
  next[len(prev)-2] = prev[len(prev)-2]
  return next
}

func extrapolate(pots []byte, rules map[string]byte, loop1 int, loop2 int) []byte {
  remaining := 50000000000 - loop2
  fmt.Println("fast forwarding", remaining, "generations")
  skiploops := remaining / (loop2 - loop1)
  fmt.Println("skipping", skiploops, "loops")
  makeup := remaining % (loop2 - loop1)
  fmt.Println("making up", makeup, "loops")
  for i := 0; i < makeup; i++ {
    pots = generation(pots, rules)
  }
  fmt.Println(string(pots), score(pots))
  return pots
}

func process(istate string, rules map[string]byte) []byte {
  history := make(map[string]int)
  pots := initialize(istate)
  history[string(pots)] = 0
  fmt.Println(0, string(pots))
  for i := 1; i <= 100000; i++ {
    pots = generation(pots, rules)
    if val, ok := history[string(pots)]; ok {
      fmt.Println("REPEAT", val, i)
      return extrapolate(pots, rules, val, i)
    } else {
      history[string(pots)] = i
    }
    fmt.Println(i, string(pots), score(pots))
  }
  return pots
}

func score(pots []byte) int {
  result := 0
  for i := 0; i < len(pots); i++ {
    if pots[i] == '#' {
      result += i-40
    }
  }
  return result
}

func main() {
  istate, transitions := readInput()
  pots := process(istate, transitions)
  result := score(pots)
  fmt.Println(result)
}
