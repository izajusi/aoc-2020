package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadStrings(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	var strs []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}

	return strs
}

func ReadInts(path string) []int {
	f, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	var ints []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Panic(err)
		}

		ints = append(ints, i)
	}

	return ints
}
