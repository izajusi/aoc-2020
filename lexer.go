package aoc

import (
	"fmt"
	"strings"
)

type ItemType int

const ItError ItemType = -1

type Item struct {
	Ty  ItemType
	Val string
}

// End of token.
const LexerEOT string = ""

type Lexer struct {
	Start    int
	Pos      int
	Width    int
	InputTok []string

	delimiter string
	items     []Item
}

type StateF func(*Lexer) StateF

func NewLexer(input string, delimiter string) *Lexer {
	return &Lexer{
		InputTok:  strings.Split(input, delimiter),
		delimiter: delimiter,
	}
}

func (l *Lexer) Run(initState StateF) []Item {
	for state := initState; state != nil; {
		state = state(l)
	}

	return l.items
}

func (l *Lexer) Emit(ty ItemType) {
	l.items = append(l.items, Item{
		Ty:  ty,
		Val: strings.Join(l.InputTok[l.Start:l.Pos], l.delimiter),
	})

	l.Start = l.Pos
}

func (l *Lexer) Next() string {
	if l.Pos >= len(l.InputTok) {
		l.Width = 0
		return LexerEOT
	}

	l.Width = 1
	l.Pos += l.Width
	return l.InputTok[l.Pos]
}

func (l *Lexer) Ignore() {
	l.Start = l.Pos
}

func (l *Lexer) Backup() {
	l.Pos -= l.Width
}

func (l *Lexer) Peek() string {
	str := l.Next()
	l.Backup()
	return str
}

func (l *Lexer) Errorf(format string, args ...interface{}) StateF {
	l.items = append(l.items, Item{
		Ty:  ItError,
		Val: fmt.Sprintf(format, args...),
	})
	return nil
}
