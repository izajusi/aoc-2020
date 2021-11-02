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

	var cnt int
	strs := aoc.ReadStrings(*path)
	bg := getGraph(strs)
	if *v2 {
		cnt = getContainedCnt(bg, "shiny gold")
	} else {
		cnt = getContainingCnt(bg, "shiny gold")
	}

	fmt.Println(cnt)
}

type bagRule struct {
	containingBag string
	containedBags map[string]int
}

type bagGraph map[string]*bagNode

type bagNode struct {
	bagType     string
	contains    bagRelations
	containedBy bagRelations
}

type bagRelations map[string]bagRelation

type bagRelation struct {
	node *bagNode
	val  int
}

func getGraph(strs []string) bagGraph {
	bg := make(bagGraph)
	p := newParser()

	getOrInitNode := func(ty string) *bagNode {
		bn, ok := bg[ty]
		if !ok {
			bn = &bagNode{
				bagType:     ty,
				contains:    make(bagRelations),
				containedBy: make(bagRelations),
			}
			bg[ty] = bn
		}

		return bn
	}

	for _, str := range strs {
		l := aoc.NewLexer(str, " ")
		items := l.Run(lexContainingBagType)
		br, err := p.parse(items)
		if err != nil {
			log.Panic(err)
		}

		bn := getOrInitNode(br.containingBag)
		for cty, val := range br.containedBags {
			cbn := getOrInitNode(cty)
			bn.contains[cty] = bagRelation{node: cbn, val: val}
			cbn.containedBy[bn.bagType] = bagRelation{node: bn, val: val}
		}

	}

	return bg
}

func getContainingCnt(bg bagGraph, ty string) int {
	walked := make(map[string]struct{})

	var walk func(ty string)
	walk = func(_ty string) {
		for _, bn := range bg[_ty].containedBy {
			cty := bn.node.bagType
			if _, ok := walked[cty]; ok {
				continue
			}

			walked[cty] = struct{}{}
			walk(cty)
		}
	}

	walk(ty)
	return len(walked)
}

func getContainedCnt(bg bagGraph, ty string) int {
	var cnt int

	var walk func(ty string, prevCnt int)
	walk = func(_ty string, prevCnt int) {
		for _, bn := range bg[_ty].contains {
			cty := bn.node.bagType
			curCnt := prevCnt * bn.val
			cnt += curCnt
			walk(cty, curCnt)
		}
	}

	walk(ty, 1)
	return cnt
}
