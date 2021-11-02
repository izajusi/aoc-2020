package main

import (
	"strings"

	"github.com/izajusi/aoc-2020"
)

const (
	itContainedBag aoc.ItemType = iota
	itContainingBag
)

func lexContainingBagType(l *aoc.Lexer) aoc.StateF {
	for {
		switch str := l.Next(); {
		case str == aoc.LexerEOT:
			return l.Errorf("reached EOT without cointaining bag keyword")
		case strings.HasPrefix(str, "bag"):
			return lexContainingBagKW
		}
	}
}

func lexContainingBagKW(l *aoc.Lexer) aoc.StateF {
	l.Pos++
	l.Emit(itContainingBag)
	return lexContain
}

func lexContain(l *aoc.Lexer) aoc.StateF {
	if l.InputTok[l.Pos] != "contain" {
		return l.Errorf(`expected "contain" keyword after containing bag`)
	}

	l.Pos++
	l.Ignore()

	return lexContainedBagType
}

func lexContainedBagType(l *aoc.Lexer) aoc.StateF {
	for {
		switch str := l.Next(); {
		case str == aoc.LexerEOT:
			return l.Errorf("reached EOT without cointained bag keyword")
		case strings.HasPrefix(str, "bag"):
			return lexContainedBagKW
		}
	}
}

func lexContainedBagKW(l *aoc.Lexer) aoc.StateF {
	l.Pos++
	l.Emit(itContainedBag)
	if l.Next() == aoc.LexerEOT {
		return nil
	}

	return lexContainedBagType
}
