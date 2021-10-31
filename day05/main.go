package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/izajusi/aoc-2020"
)

func main() {
	var path = flag.String("p", "", "input file path")
	var v2 = flag.Bool("v2", false, "return answer for part 2")
	flag.Parse()

	var id int
	strs := aoc.ReadStrings(*path)
	if *v2 {
		id = getMissingID(strs)
	} else {
		id = getMaxID(strs)
	}

	fmt.Println(id)
}

func getMaxID(strs []string) int {
	var max int

	for _, str := range strs {
		st := parse(str)
		if st.id > max {
			max = st.id
		}
	}

	return max
}

func getMissingID(strs []string) int {
	seatMap := make(map[int]*seat)

	for _, str := range strs {
		st := parse(str)
		seatMap[st.id] = &st
	}

	const (
		minRow int = 0
		maxRow int = 127
		minCol int = 0
		maxCol int = 7
	)

	for row := minRow; row < maxRow; row++ {
		for col := minCol; col < maxCol; col++ {
			id := calcID(row, col)
			before, after := seatMap[id-1], seatMap[id+1]
			if seatMap[id] == nil && before != nil && after != nil {
				return id
			}
		}
	}

	return -1
}

type seat struct {
	id  int
	row int
	col int
}

func parse(str string) seat {
	rowPart := str[:7]
	colPart := str[7:]

	st := seat{
		row: strToInt(rowPart, 'B', 'F'),
		col: strToInt(colPart, 'R', 'L'),
	}

	st.id = calcID(st.row, st.col)
	return st
}

func calcID(row int, col int) int {
	return (row * 8) + col
}

func strToInt(str string, high rune, low rune) int {
	var val int
	var mult int = 1

	runes := []rune(str)
	for i := len(str) - 1; i >= 0; i-- {
		if runes[i] == high {
			val += mult
		} else if runes[i] != low {
			log.Panicf("unexpected rune: %v", runes[i])
		}

		mult *= 2
	}

	return val
}
