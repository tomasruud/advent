package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	ps := parse(input)

	fmt.Println(ps.valid())
	fmt.Println(ps.validAtToboggan())
}

func parse(in string) list {
	r := regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<char>[A-Za-z]): (?P<pass>[A-Za-z]+)`)

	var ps list
	for _, l := range strings.Split(in, "\n") {
		match := r.FindStringSubmatch(l)

		sub := make(map[string]string)
		for i, name := range r.SubexpNames() {
			if i != 0 && name != "" && i < len(match) {
				sub[name] = match[i]
			}
		}

		min, _ := strconv.Atoi(sub["min"])
		max, _ := strconv.Atoi(sub["max"])

		ps = append(ps, pw{
			min:  min,
			max:  max,
			char: sub["char"],
			pass: sub["pass"],
		})
	}

	return ps
}

type list []pw

type pw struct {
	min  int
	max  int
	char string
	pass string
}

func (p pw) valid() bool {
	c := strings.Count(p.pass, p.char)

	if c >= p.min && c <= p.max {
		return true
	}

	return false
}

func (p pw) validAtToboggan() bool {
	l := len(p.pass)
	first := p.min - 1
	second := p.max - 1

	matched := false

	if first >= 0 && first < l && string(p.pass[first]) == p.char {
		matched = true
	}

	if second >= 0 && second < l && string(p.pass[second]) == p.char {
		matched = !matched
	}

	return matched
}

func (ps list) valid() int {
	n := 0
	for _, p := range ps {
		if p.valid() {
			n++
		}
	}
	return n
}

func (ps list) validAtToboggan() int {
	n := 0
	for _, p := range ps {
		if p.validAtToboggan() {
			n++
		}
	}
	return n
}
