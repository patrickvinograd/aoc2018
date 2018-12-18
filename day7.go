package main

import (
	"bufio"
	"fmt"
	"os"
  "sort"
  "strings"
)

type Vertex struct {
	From string
	To   string
}

type Coord struct {
	Prox int
	X    int
	Y    int
}

func deleteV(vPtr *[]Vertex, i int) {
  v := *vPtr
  copy(v[i:], v[i+1:])
  v[len(v)-1] = Vertex{} // or the zero value of T
  *vPtr = v[:len(v)-1]
}

func deleteS(vPtr *[]string, i int) {
  v := *vPtr
  copy(v[i:], v[i+1:])
  v[len(v)-1] = "" // or the zero value of T
  *vPtr = v[:len(v)-1]
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func mapkeys(in map[string]struct{}) []string {
  keys := make([]string, len(in))

  i := 0
  for k := range in {
    keys[i] = k
    i++
  }
  return keys
}

func process(input []Vertex) []string {
  result := make([]string, 0)
  readylist := make([]string, 0)

  keys := make(map[string]struct{})
  for _, v := range input {
    keys[v.From] = struct{}{}
    keys[v.To] = struct{}{}
  }

  for {
    for k, _ := range keys {
      anyTo := false
      for _, v := range input {
        if v.To == k {
          anyTo = true
        }
      }
      if anyTo == false && !contains(readylist, k) {
        readylist = append(readylist, k)
      }
    }
    sort.Strings(readylist)
    //fmt.Println("readylist", readylist)
    next := readylist[0]
    deleteS(&readylist, 0)
    delete(keys, next)
    result = append(result, next)
    for i := 0; i < len(input); {
      v := input[i]
      if v.From == next {
        //fmt.Println("deleting", v)
        deleteV(&input, i)
      } else {
        //fmt.Println("skipping", v)
        i++
      }
    }
    //fmt.Println("result", result)
    //fmt.Println("input", input)
    if len(keys) == 0 {
      return result
    }
  }
}

func readInput() []Vertex {
	result := make([]Vertex, 0)
	reader := bufio.NewReader(os.Stdin)
	bytes, _, err := reader.ReadLine()
	for err == nil {
		in := string(bytes)
		v := Vertex{}
		fmt.Sscanf(in, "Step %1s must be finished before step %1s can begin.", &v.From, &v.To)
		result = append(result, v)
		bytes, _, err = reader.ReadLine()
	}
	return result
}

func main() {
	vertices := readInput()
	fmt.Println(vertices)
  result := process(vertices)
  fmt.Println(strings.Join(result, ""))
}
