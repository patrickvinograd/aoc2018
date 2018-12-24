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
  for i := 0; i < 40; i++ {
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
    //fmt.Println(rule,  "=>", string(out))
    next[i] = out
  }
  next[0] = prev[0]
  next[1] = prev[1]
  next[len(prev)-1] = prev[len(prev)-1]
  next[len(prev)-2] = prev[len(prev)-2]
  return next
}

func process(istate string, rules map[string]byte) []byte {
  pots := initialize(istate)
  fmt.Println(0, string(pots))
  for i := 1; i <= 20; i++ {
    pots = generation(pots, rules)
    fmt.Println(i, string(pots))
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
