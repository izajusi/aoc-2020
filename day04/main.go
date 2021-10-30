package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var path = flag.String("p", "", "input file path")
	var v2 = flag.Bool("v2", false, "return answer for part 2")
	flag.Parse()

	var validF func(passport) bool
	if *v2 {
		validF = validV2
	} else {
		validF = valid
	}

	var cnt int
	for _, pass := range getPassports(*path) {
		if validF(pass) {
			cnt++
		}
	}

	fmt.Println(cnt)
}

type passport map[string]string

func getPassports(path string) []passport {
	f, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	parsePassport := func(pass passport, line string) {
		strs := strings.Split(line, " ")
		for _, str := range strs {
			kv := strings.Split(str, ":")
			pass[kv[0]] = kv[1]
		}
	}

	var passports []passport
	scanner := bufio.NewScanner(f)
	pass := make(passport)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, pass)
			pass = make(passport)
			continue
		}

		parsePassport(pass, line)
	}

	// Flush the last passport.
	passports = append(passports, pass)

	return passports
}

type fieldValidFunc func(string) bool

var reqFields = map[string]fieldValidFunc{
	"byr": validRangeGen(1920, 2002),
	"iyr": validRangeGen(2010, 2020),
	"eyr": validRangeGen(2020, 2030),
	"hgt": validHeightGen(),
	"hcl": validRegexGen(`^#[a-f0-9]{6}$`),
	"ecl": validRegexGen(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
	"pid": validRegexGen(`^\d{9}$`),
}

func valid(pass passport) bool {
	for field := range reqFields {
		if _, ok := pass[field]; !ok {
			return false
		}
	}

	return true
}

func validV2(pass passport) bool {
	for field, validF := range reqFields {
		if val, ok := pass[field]; !ok || !validF(val) {
			return false
		}
	}

	return true
}

func validRangeGen(from int, to int) fieldValidFunc {
	return func(s string) bool {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}

		return i >= from && i <= to
	}
}

func validHeightGen() fieldValidFunc {
	re := regexp.MustCompile(`^(\d+)(in|cm)$`)

	return func(s string) bool {
		matches := re.FindStringSubmatch(s)
		if len(matches) != 3 {
			return false
		}

		val := matches[1]
		unit := matches[2]

		cmF := validRangeGen(150, 193)
		inF := validRangeGen(59, 76)
		if unit == "cm" {
			return cmF(val)
		} else if unit == "in" {
			return inF(val)
		} else {
			return false
		}
	}
}

func validRegexGen(exp string) fieldValidFunc {
	re := regexp.MustCompile(exp)

	return func(s string) bool {
		return re.Match([]byte(s))
	}
}
