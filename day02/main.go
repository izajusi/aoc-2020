package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/izajusi/aoc-2020"
)

func main() {
	var path = flag.String("p", "", "input file path")
	var v2 = flag.Bool("v2", false, "return answer for part 2")
	flag.Parse()

	var validF func(policy, string) bool
	if *v2 {
		validF = validV2
	} else {
		validF = valid
	}

	var cnt int
	strs := aoc.ReadStrings(*path)
	for _, str := range strs {
		pol, pass := parse(str)
		if validF(pol, pass) {
			cnt++
		}
	}

	fmt.Println(cnt)
}

type policy struct {
	word rune
	min  int
	max  int
}

func parse(str string) (policy, string) {
	var err error
	var pol policy

	strs := strings.Split(str, " ")

	// First part: "{min}-{max}"
	first := strings.Split(strs[0], "-")
	if pol.min, err = strconv.Atoi(first[0]); err != nil {
		log.Panic(err)
	}
	if pol.max, err = strconv.Atoi(first[1]); err != nil {
		log.Panic(err)
	}

	// Second part: "{word}:"
	pol.word = []rune(strs[1])[0]

	// Third part: "{password}"
	pass := strs[2]

	return pol, pass
}

func valid(pol policy, pass string) bool {
	var cnt int

	for _, c := range pass {
		if c == pol.word {
			cnt++
		}
	}

	return cnt >= pol.min && cnt <= pol.max
}

func validV2(pol policy, pass string) bool {
	first := pol.min - 1
	second := pol.max - 1
	runes := []rune(pass)

	// XOR: https://stackoverflow.com/a/23025720
	return (runes[first] == pol.word) != (runes[second] == pol.word)
}
