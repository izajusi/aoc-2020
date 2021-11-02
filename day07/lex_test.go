package main

import (
	"testing"

	"github.com/izajusi/aoc-2020"
	"github.com/stretchr/testify/assert"
)

func TestLexBag(t *testing.T) {
	input := "muted gold bags contain 1 wavy red bag, 3 mirrored violet bags, 5 bright gold bags, 5 plaid white bags."
	lexer := aoc.NewLexer(input, " ")
	items := lexer.Run(lexContainingBagType)

	assert.Len(t, items, 5)
	assert.Subset(t, items, []aoc.Item{
		{Ty: itContainingBag, Val: "muted gold bags"},
		{Ty: itContainedBag, Val: "1 wavy red bag,"},
		{Ty: itContainedBag, Val: "3 mirrored violet bags,"},
		{Ty: itContainedBag, Val: "5 bright gold bags,"},
		{Ty: itContainedBag, Val: "5 plaid white bags."},
	})
}

func TestLexEmptyBag(t *testing.T) {
	input := "drab silver bags contain no other bags."
	lexer := aoc.NewLexer(input, " ")
	items := lexer.Run(lexContainingBagType)

	assert.Len(t, items, 2)
	assert.Subset(t, items, []aoc.Item{
		{Ty: itContainingBag, Val: "drab silver bags"},
		{Ty: itContainedBag, Val: "no other bags."},
	})
}
