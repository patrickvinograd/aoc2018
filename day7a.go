package main

import (
	"bufio"
	"fmt"
	"os"
  "sort"
  "strings"
)

const offset = 61
const elfCount = 5

type Vertex struct {
	From string
	To   string
}

type ElfJob struct {
  Id int
  Task string
  EndTime int
}

func makeTimings() map[string]int {
  counts := make(map[string]int)
  start := int('A')
  for i := 0; i < 26; i++ {
    s := string(start + i)
    c := int(i + offset)
    counts[s] = c
  }
  return counts
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

func nextElf(elves []*ElfJob, time int) (int, int) {
  id := -1
  endtime := 10000000
  for i := 0; i < len(elves); i++ {
    if elves[i].EndTime < endtime {
      id = elves[i].Id
      endtime = elves[i].EndTime
    }
  }
  return id, endtime
}


func process(input []Vertex) ([]string, int) {
  result := make([]string, 0)
  readylist := make([]string, 0)

  keys := make(map[string]struct{})
  startedkeys := make(map[string]bool)
  for _, v := range input {
    keys[v.From] = struct{}{}
    keys[v.To] = struct{}{}
    startedkeys[v.From] = false
    startedkeys[v.To] = false
  }
  timings := makeTimings()

  time := -1
  workers := make([]*ElfJob, 0)
  for i := 0; i < elfCount; i++ {
    workers = append(workers, &ElfJob{Id: i, EndTime: -1})
  }

  for {
    // check for any tasks wrapping up
    for i:= 0; i < len(workers); i++ {
      e := workers[i]
      if e.Task != "" && e.EndTime == time {
        //fmt.Println("DoneTask", e)
        result = append(result, e.Task)
        delete(keys, e.Task)
        //delete inbound vertexes for completed task
        for j := 0; j < len(input); {
          v := input[j]
          if v.From == e.Task {
            deleteV(&input, j)
          } else {
            j++
          }
        }
        e.Task = ""
      }
    }
    //fmt.Println("keys", time, keys)
    //push any newly freed tasks onto the readylist
    for k, _ := range keys {
      anyTo := false
      for _, v := range input {
        if v.To == k {
          anyTo = true
        }
      }
      if anyTo == false && !contains(readylist, k) && startedkeys[k] != true {
        readylist = append(readylist, k)
        startedkeys[k] = true
      }
    }
    sort.Strings(readylist)
    //fmt.Println(time, readylist)

    //find any free elves, give them any tasks off readylist
    for i:= 0; i < len(workers); i++ {
      e := workers[i]
      if e.Task == "" && len(readylist) > 0 {
        nextTask := readylist[0]
        deleteS(&readylist, 0)
        e.Task = nextTask
        e.EndTime = time + timings[nextTask]
        fmt.Println("AssignTask", e)
      }
    }
    //fmt.Println(time, readylist)

    time++
    if len(keys) == 0 {
      return result, time
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
  result, time := process(vertices)
  fmt.Println(strings.Join(result, ""))
  fmt.Println(time)
}
