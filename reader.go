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

func ReadStringArrs(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	strArrs := make([][]string, 0)
	strs := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			strArrs = append(strArrs, strs)
			strs = make([]string, 0)
			continue
		}

		strs = append(strs, line)
	}

	// Flush the last array.
	strArrs = append(strArrs, strs)

	return strArrs
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
