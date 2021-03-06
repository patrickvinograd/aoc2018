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

func initOps() []Op {
  result := []Op{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
  return result
}

type Sample struct {
  Before []int
  Code int
  A int
  B int
  C int
  After []int
}

func readSamples() []Sample {
  reader := bufio.NewReader(os.Stdin)
  result := make([]Sample, 0)
  bytes, _, err := reader.ReadLine()
  for err == nil {
    s := Sample{Before: make([]int, 4), After: make([]int, 4)}
    in := string(bytes)
    fmt.Sscanf(in, "Before: [%d, %d, %d, %d]", &s.Before[0], &s.Before[1], &s.Before[2], &s.Before[3])
    bytes, _, err = reader.ReadLine()
    in = string(bytes)
    fmt.Sscanf(in, "%d %d %d %d", &s.Code, &s.A, &s.B, &s.C)
    bytes, _, err = reader.ReadLine()
    in = string(bytes)
    fmt.Sscanf(in, "After: [%d, %d, %d, %d]", &s.After[0], &s.After[1], &s.After[2], &s.After[3])
    result = append(result, s)
    bytes, _, err = reader.ReadLine()
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

func try(sample Sample, ops []Op) []int {
  result := make([]int, 0)
  for _, op := range ops {
    //fmt.Println("Checking", sample.Code, "for", sample.After)
    output := op(sample.A, sample.B, sample.C, sample.Before)
    //fmt.Println(output)
    if Equal(output, sample.After) {
      result = append(result, sample.Code)
    }
  }
  return result
}

func main() {
  samples := readSamples()
  ops := initOps()
  fmt.Println(samples)
  result := 0
  for _, s := range samples {
    matches := try(s, ops)
    if len(matches) >= 3 {
      result++
    }
  }
  fmt.Println(result)
}
