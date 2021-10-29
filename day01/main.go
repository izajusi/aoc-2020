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

	var answer int
	ints := aoc.ReadInts(*path)
	if *v2 {
		answer = findThrees(ints)
	} else {
		answer = findPair(ints)
	}

	fmt.Println(answer)
}

func findPair(ints []int) int {
	for i, first := range ints {
		for _, second := range ints[i+1:] {
			if first+second == 2020 {
				return first * second
			}
		}
	}

	return -1
}

func findThrees(ints []int) int {
	for i, first := range ints {
		for j, second := range ints[i+1:] {
			if first+second >= 2020 {
				continue
			}

			for _, third := range ints[j+1:] {
				if first+second+third == 2020 {
					return first * second * third
				}
			}
		}
	}

	return -1
}
