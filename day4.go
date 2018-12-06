package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Range struct {
	Start int
	End   int
}

type GuardRecord struct {
	Id     int
	Ranges []Range
}

func readInput() []string {
	result := make([]string, 0)
	//var in string
	reader := bufio.NewReader(os.Stdin)
	bytes, _, err := reader.ReadLine()
	//lines, _ := fmt.Scanln(&in)
	for err == nil {
		in := string(bytes)
		result = append(result, in)
		bytes, _, err = reader.ReadLine()
	}
	sort.Strings(result)
	return result
}

func parse(lines []string) []*GuardRecord {
	result := make([]*GuardRecord, 0)
	var currentGuard *GuardRecord
	var currentStart int
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		//fmt.Println(line)
		if strings.HasSuffix(line, "begins shift") {
			var y, m, d, hh, mm, id int
			var g, b, s string
			// [1518-07-03 23:58] Guard #2437 begins shift
			fmt.Sscanf(line, "[%d-%d-%d %d:%d] %v #%d %v %v",
				&y, &m, &d, &hh, &mm, &g, &id, &b, &s)
			newGuard := GuardRecord{id, make([]Range, 0)}
			result = append(result, &newGuard)
			currentGuard = &newGuard
			//fmt.Println(id)
		}
		if strings.HasSuffix(line, "falls asleep") {
			var y, m, d, hh, mm int
			var w1, w2 string
			fmt.Sscanf(line, "[%d-%d-%d %d:%d] %v #%d %v %v",
				&y, &m, &d, &hh, &mm, &w1, &w2)
			currentStart = mm
			//fmt.Println(currentStart)
		}
		if strings.HasSuffix(line, "wakes up") {
			var y, m, d, hh, mm int
			var w1, w2 string
			fmt.Sscanf(line, "[%d-%d-%d %d:%d] %v #%d %v %v",
				&y, &m, &d, &hh, &mm, &w1, &w2)
			currentGuard.Ranges = append(currentGuard.Ranges, Range{currentStart, mm})
			//fmt.Println(mm)
		}
		//fmt.Println(currentGuard)
	}
	return result
}

func worstGuard(recs []*GuardRecord) int {
	sleeps := make(map[int]int)
	for i := 0; i < len(recs); i++ {
		rec := recs[i]
		total := 0
		for j := 0; j < len(rec.Ranges); j++ {
			r := rec.Ranges[j]
			total = total + (r.End - r.Start)
		}
		sleeps[rec.Id] = sleeps[rec.Id] + total
	}
	//fmt.Println(sleeps)
	result := 0
	max := 0
	for k, v := range sleeps {
		if v > max {
			result = k
			max = v
		}
	}
	return result
}

func worstMinute(recs []*GuardRecord, wg int) int {
	mins := make(map[int]int)
	for _, rec := range recs {
		if rec.Id == wg {
			//fmt.Println(rec)
			for _, rang := range rec.Ranges {
				for i := rang.Start; i <= rang.End; i++ {
					mins[i] = mins[i] + 1
				}
			}
		}
	}
	result := 0
	max := 0
	for k, v := range mins {
		if v > max {
			result = k
			max = v
		}
	}
	return result
}

func main() {
	values := readInput()
	records := parse(values)
	//fmt.Println(records)
	worst := worstGuard(records)
	fmt.Println(worst)
	worstMin := worstMinute(records, worst)
	fmt.Println(worstMin)
	fmt.Println((worst * worstMin))
}
