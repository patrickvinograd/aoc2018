package main

import (
	"bufio"
  "io/ioutil"
	"fmt"
	"os"
)

func readInput() []byte {
	reader := bufio.NewReader(os.Stdin)
  bytes, _ := ioutil.ReadAll(reader)
  if bytes[len(bytes) -1] == 10 {
    bytes = bytes[0:len(bytes)-1]
  }
  return bytes
}

func delete2(aPtr *[]byte, i int) {
  a := *aPtr
  copy(a[i:], a[i+2:])
  a[len(a)-1] = 0 // or the zero value of T
  a[len(a)-2] = 0 // or the zero value of T
  *aPtr = a[:len(a)-2]
}

func reduce(inputPtr *[]byte) int {
  input := *inputPtr
  i := 0
  //l := len(input)
  for {
//    fmt.Println(i, input)
    if i == len(input) - 1 {
      return len(input)
    }
    a, b := input[i], input[i+1]
    if (a - b == 32 || b - a == 32) {
      delete2(&input, i)
      if (i > 0) {
        i = i - 1
      }
    } else {
      i++
    }
  }
}

func main() {
	values := readInput()
  //fmt.Println(values)
  result := reduce(&values)
  //fmt.Println(string(values))
  fmt.Println(len(values))
  fmt.Println(result)
}
