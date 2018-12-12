package main

import (
	"bufio"
  "io/ioutil"
	"fmt"
	"os"
)

var points = [] byte("AaBbCcDd")

func readInput() []byte {
	reader := bufio.NewReader(os.Stdin)
  bytes, _ := ioutil.ReadAll(reader)
  if bytes[len(bytes) -1] == 10 {
    bytes = bytes[0:len(bytes)-1]
  }
  return bytes
}

func delete1(aPtr *[]byte, i int) {
  a := *aPtr
  copy(a[i:], a[i+1:])
  a[len(a)-1] = 0 // or the zero value of T
  *aPtr = a[:len(a)-1]
}

func delete2(aPtr *[]byte, i int) {
  a := *aPtr
  copy(a[i:], a[i+2:])
  a[len(a)-1] = 0 // or the zero value of T
  a[len(a)-2] = 0 // or the zero value of T
  *aPtr = a[:len(a)-2]
}

func reduce(inputPtr *[]byte, rem1 byte, rem2 byte) int {
  input := *inputPtr
  i := 0
  for {
    //fmt.Println(i, input)
    if i == len(input) {
      *inputPtr = input
      return len(input)
    }
    //a, b := input[i], input[i+1]
    if a := input[i]; (a == rem1 || a == rem2) {
      delete1(&input, i)
      if (i > 0) {
        i = i - 1
      }
    } else if (i < len(input) -1 && (input[i] - input[i+1] == 32 || input[i+1] - input[i] == 32)) {
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
  best := 1000000
  var i byte = 65
  for ; i <= 90; i++ {
    a := make([]byte, len(values))
    copy(a, values)
    aresult := reduce(&a, i, i+32)
    fmt.Println(i, aresult)
    if aresult < best {
      best = aresult
    }
  }
  fmt.Println(best)
}
