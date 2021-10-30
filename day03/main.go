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
	strs := aoc.ReadStrings(*path)
	initWidth := len(strs[0])
	if *v2 {
		cnt = traverse(strs, slopeGen(1, 1, initWidth)) *
			traverse(strs, slopeGen(3, 1, initWidth)) *
			traverse(strs, slopeGen(5, 1, initWidth)) *
			traverse(strs, slopeGen(7, 1, initWidth)) *
			traverse(strs, slopeGen(1, 2, initWidth))
	} else {
		cnt = traverse(strs, slopeGen(3, 1, initWidth))
	}

	fmt.Println(cnt)
}

type slopeFunc func(int, int) (int, int)

func slopeGen(xMov, yMov int, initWidth int) slopeFunc {
	return func(x, y int) (int, int) {
		nx := x + xMov
		ny := y + yMov

		if nx > (initWidth - 1) {
			nx = nx - initWidth
		}

		return nx, ny
	}
}

func traverse(strs []string, slopeF slopeFunc) int {
	var cnt int
	height := len(strs)

	for x, y := 0, 0; y < height; x, y = slopeF(x, y) {
		runes := []rune(strs[y])
		if runes[x] == '#' {
			cnt++
		}
	}

	return cnt
}
