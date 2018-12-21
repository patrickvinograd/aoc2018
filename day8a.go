package main

import (
	"fmt"
)


type Node struct {
	ChildCount int
  MetaCount int
  ChildValues []int
  Metadata []int
  Value int
}

func delete(vPtr *[]int, i int) {
  v := *vPtr
  copy(v[i:], v[i+1:])
  v[len(v)-1] = 0 // or the zero value of T
  *vPtr = v[:len(v)-1]
}

func deleteNode(vPtr *[]Node, i int) {
  v := *vPtr
  copy(v[i:], v[i+1:])
  v[len(v)-1] = Node{} // or the zero value of T
  *vPtr = v[:len(v)-1]
}

func readInput() []int {
        result := make([]int, 0)
        var token int
        numread, _ := fmt.Scan(&token)
        for numread > 0 {
                result = append(result, token)
                numread, _ = fmt.Scan(&token)
        }
        return result
}

func takeNode(accumPtr *[]Node, stream []int, index int) (int, int) {
  accum := *accumPtr
  currentNode := Node{Metadata: make([]int, 0), ChildValues: make([]int, 0)}
  currentNode.ChildCount = stream[index]
  index++
  currentNode.MetaCount = stream[index]
  index++
  for c := 0; c < currentNode.ChildCount; c++ {
    var cval int
    index, cval = takeNode(&accum, stream, index)
    currentNode.ChildValues = append(currentNode.ChildValues, cval)
  }
  for m := 0; m < currentNode.MetaCount; m++ {
    currentNode.Metadata = append(currentNode.Metadata, stream[index])
    index++
  }
  if currentNode.ChildCount == 0 {
    value := 0
    for i := 0; i < len(currentNode.Metadata); i++ {
      value += currentNode.Metadata[i]
    }
    currentNode.Value = value
  } else {
    value := 0
    for i := 0; i < len(currentNode.Metadata); i++ {
      m := currentNode.Metadata[i] - 1
      if m < currentNode.ChildCount {
        value += currentNode.ChildValues[m]
      }
    }
    currentNode.Value = value
  }
  accum = append(accum, currentNode)
  *accumPtr = accum
  return index, currentNode.Value
}

func process(stream []int) ([]Node, int) {
  total := 0
  result := make([]Node, 0)
  for i := 0; i < len(stream); {
    i, total = takeNode(&result, stream, i)
  }
  return result, total
}

func mdTotal(nodes []Node) int {
  result := 0
  for i := 0; i < len(nodes); i++ {
    n := nodes[i]
    for j := 0; j < len(n.Metadata); j++ {
      result += n.Metadata[j]
    }
  }
  return result
}

func main() {
	nodestream := readInput()
  fmt.Println(len(nodestream))
  nodes, total := process(nodestream)
  fmt.Println(nodes)
  //result := mdTotal(nodes)
  fmt.Println(total)
}
