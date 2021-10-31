package main

import (
	"flag"
	"fmt"

	"github.com/izajusi/aoc-2020"
)

func main() {
	var path = flag.String("p", "", "input file path")
	var v2 = flag.Bool("v2", false, "return answer for part 2")
	flag.Parse()

	var cnt int
	strArrs := aoc.ReadStringArrs(*path)
	for _, strs := range strArrs {
		ans := parseGrpAnswers(strs)

		if *v2 {
			cnt += unanimousCnt(ans)
		} else {
			cnt += len(ans.answers)
		}

	}

	fmt.Println(cnt)
}

type grpAnswers struct {
	groupSize int
	answers   map[rune]int
}

func parseGrpAnswers(lines []string) grpAnswers {
	var ans grpAnswers
	ans.groupSize = len(lines)
	ans.answers = make(map[rune]int)

	for _, line := range lines {
		for _, c := range line {
			ans.answers[c]++
		}
	}

	return ans
}

func unanimousCnt(ans grpAnswers) int {
	var unCnt int

	for _, cnt := range ans.answers {
		if cnt == ans.groupSize {
			unCnt++
		}
	}

	return unCnt
}
