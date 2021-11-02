package main

import (
	"testing"

	"github.com/izajusi/aoc-2020"
	"github.com/stretchr/testify/assert"
)

func TestParseBagRule(t *testing.T) {
	items := []aoc.Item{
		{Ty: itContainingBag, Val: "muted gold bags"},
		{Ty: itContainedBag, Val: "1 wavy red bag,"},
		{Ty: itContainedBag, Val: "3 mirrored violet bags,"},
		{Ty: itContainedBag, Val: "5 bright gold bags,"},
		{Ty: itContainedBag, Val: "5 plaid white bags."},
	}

	p := newParser()
	br, err := p.parse(items)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &bagRule{
			containingBag: "muted gold",
			containedBags: map[string]int{
				"wavy red":        1,
				"mirrored violet": 3,
				"bright gold":     5,
				"plaid white":     5,
			},
		}, br)
	}
}

func TestParseEmptyBagRule(t *testing.T) {
	items := []aoc.Item{
		{Ty: itContainingBag, Val: "drab silver bags"},
		{Ty: itContainedBag, Val: "no other bags."},
	}

	p := newParser()
	br, err := p.parse(items)
	if assert.NoError(t, err) {
		assert.EqualValues(t, "drab silver", br.containingBag)
		assert.Len(t, br.containedBags, 0)
	}
}
