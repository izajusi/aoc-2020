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

	var acc int
	strs := aoc.ReadStrings(*path)
	firstIns := getFirstIns(strs)
	if *v2 {
		acc = getAccAtFixLoop(firstIns)
	} else {
		acc = getAccAtLoop(firstIns)
	}

	fmt.Println(acc)
}

type instructionSet []*instruction

type instructionType int

const (
	insJmp instructionType = iota
	insAcc
	insNop
)

type instruction struct {
	insSet instructionSet
	ty     instructionType
	id     int
	val    int
}

func (i *instruction) exec(acc int, sw bool) (nextAcc int, nextIns *instruction) {
	var nextID int
	nextAcc = acc

	switch ty := i.ty; {
	case (ty == insJmp && !sw) || (ty == insNop && sw):
		nextID = i.id + i.val
	case ty == insAcc:
		nextAcc += i.val
		nextID = i.id + 1
	case (ty == insNop && !sw) || (ty == insJmp && sw):
		nextID = i.id + 1
	}

	if nextID == len(i.insSet) {
		return nextAcc, nil // Instruction set done.
	} else if nextID > len(i.insSet) {
		log.Panic("out of index")
	}

	return nextAcc, i.insSet[nextID]
}

func getFirstIns(lines []string) *instruction {
	insSet := make(instructionSet, len(lines))

	for i, line := range lines {
		strs := strings.Split(line, " ")
		if len(strs) != 2 {
			log.Panic("unexpected result when parsing")
		}

		var ty instructionType
		switch strs[0] {
		case "acc":
			ty = insAcc
		case "jmp":
			ty = insJmp
		case "nop":
			ty = insNop
		default:
			log.Panicf("unexpected instruction: %v", strs[0])
		}

		val, err := strconv.Atoi(strs[1])
		if err != nil {
			log.Panicf("unexpected value: %v", strs[1])
		}

		ins := &instruction{
			insSet: insSet,
			ty:     ty,
			id:     i,
			val:    val,
		}

		insSet[i] = ins
	}

	return insSet[0]
}

func getAccAtLoop(firstIns *instruction) (acc int) {
	insRan := make(map[int]bool)

	var ins *instruction
	for ins = firstIns; !insRan[ins.id]; acc, ins = ins.exec(acc, false) {
		insRan[ins.id] = true
	}

	return acc
}

func getAccAtFixLoop(firstIns *instruction) int {
	type key struct {
		id    int
		hasSw bool
	}
	insRan := make(map[key]bool)

	var walk func(ins *instruction, acc int, hasSw bool) (bool, int)
	walk = func(ins *instruction, acc int, hasSw bool) (bool, int) {
		if insRan[key{ins.id, hasSw}] {
			return false, 0
		}

		insRan[key{ins.id, hasSw}] = true

		if !hasSw {
			nextAcc, nextIns := ins.exec(acc, true)
			if nextIns == nil {
				return true, nextAcc
			}

			found, nextAcc := walk(nextIns, nextAcc, true)
			if found {
				return true, nextAcc
			}
		}

		nextAcc, nextIns := ins.exec(acc, false)
		if nextIns == nil {
			return true, nextAcc
		}

		return walk(nextIns, nextAcc, hasSw)
	}

	_, acc := walk(firstIns, 0, false)
	return acc
}
