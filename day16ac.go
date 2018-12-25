package main

import (
  "fmt"
  "bufio"
  "os"
)

type Op func(int, int, int, []int)[]int

func addr(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  output[c] = output[a] + output[b]
  return output
}

func addi(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  output[c] = output[a] + b
  return output
}

func mulr(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  output[c] = output[a] * output[b]
  return output
}

func muli(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  output[c] = output[a] * b
  return output
}

func banr(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  output[c] = output[a] & output[b]
  return output
}

func bani(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  output[c] = output[a] & b
  return output
}

func borr(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  output[c] = output[a] | output[b]
  return output
}

func bori(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  output[c] = output[a] | b
  return output
}

func setr(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  output[c] = output[a]
  return output
}

func seti(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  output[c] = a
  return output
}

func gtir(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  if a > output[b] {
    output[c] = 1
  } else {
    output[c] = 0
  }
  return output
}

func gtri(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  if output[a] > b {
    output[c] = 1
  } else {
    output[c] = 0
  }
  return output
}

func gtrr(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  if output[a] > output[b] {
    output[c] = 1
  } else {
    output[c] = 0
  }
  return output
}

func eqir(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  if a == output[b] {
    output[c] = 1
  } else {
    output[c] = 0
  }
  return output
}

func eqri(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  if output[a] == b {
    output[c] = 1
  } else {
    output[c] = 0
  }
  return output
}

func eqrr(a int, b int, c int, reg []int) []int {
  output := make([]int, len(reg))
  copy(output, reg)
  if output[a] == output[b] {
    output[c] = 1
  } else {
    output[c] = 0
  }
  return output
}

var opnames = []string{"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "setr", "seti", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr"}

func initOps() map[int]Op {
  opcodes := make(map[int]Op)
  opcodes[0] = borr
  opcodes[1] = seti
  opcodes[2] = mulr
  opcodes[3] = eqri
  opcodes[4] = banr
  opcodes[5] = bori
  opcodes[6] = bani
  opcodes[7] = gtri
  opcodes[8] = addr
  opcodes[9] = muli
  opcodes[10] = addi
  opcodes[11] = eqrr
  opcodes[12] = gtir
  opcodes[13] = eqir
  opcodes[14] = setr
  opcodes[15] = gtrr

  return opcodes
}

type Line struct {
  Code int
  A int
  B int
  C int
}

func readLines() []Line {
  reader := bufio.NewReader(os.Stdin)
  result := make([]Line, 0)
  bytes, _, err := reader.ReadLine()
  for err == nil {
    l := Line{}
    in := string(bytes)
    fmt.Sscanf(in, "%d %d %d %d", &l.Code, &l.A, &l.B, &l.C)
    result = append(result, l)
    bytes, _, err = reader.ReadLine()
  }
  return result
}

func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func process(lines []Line, opcodes map[int]Op) []int {
  reg := []int{0, 0, 0, 0}
  for _, line := range lines {
    op := opcodes[line.Code]
    reg = op(line.A, line.B, line.C, reg)
  }
  return reg
}

func main() {
  lines := readLines()
  opcodes := initOps()
  regs := process(lines, opcodes)
  fmt.Println(regs[0])
}
