package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/izajusi/aoc-2020"
)

type parser struct {
	reContainingBag *regexp.Regexp
	reContainedBag  *regexp.Regexp
}

func newParser() *parser {
	return &parser{
		reContainingBag: regexp.MustCompile(`([a-z ]+) bags`),
		reContainedBag:  regexp.MustCompile(`(\d+) ([a-z ]+) bags?`),
	}
}

func (p *parser) parse(items []aoc.Item) (*bagRule, error) {
	var br bagRule
	br.containedBags = make(map[string]int)

	for _, item := range items {
		switch item.Ty {
		case itContainingBag:
			matches := p.reContainingBag.FindStringSubmatch(item.Val)
			if len(matches) != 2 {
				return nil, fmt.Errorf("failed to parse containing bag: %q. unexpected matches: %v",
					item.Val, len(matches))
			}

			br.containingBag = matches[1]
		case itContainedBag:
			if strings.HasPrefix(item.Val, "no other bag") {
				continue
			}

			matches := p.reContainedBag.FindStringSubmatch(item.Val)
			if len(matches) != 3 {
				return nil, fmt.Errorf("failed to parse contained bag: %q. unexpected matches: %v",
					item.Val, len(matches))
			}

			count, err := strconv.Atoi(matches[1])
			if err != nil {
				return nil, fmt.Errorf("failed to parse contained bag: %w", err)
			}

			br.containedBags[matches[2]] = count
		case aoc.ItError:
			return nil, fmt.Errorf("error when parsing: %v", item.Val)
		}
	}

	if br.containingBag == "" {
		return nil, errors.New("no containing bag found")
	}

	return &br, nil
}
