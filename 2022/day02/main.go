package main

import (
	"fmt"
	"strings"
)

func main() {
	g, err := parsePart1(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("my score", g.score())

	g2, err := parsePart2(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("my actual score", g2.score())
}

func parsePart1(in string) (guide, error) {
	var out guide
	for _, ln := range strings.Split(in, "\n") {
		ln = strings.TrimSpace(ln)

		parts := strings.Split(ln, " ")
		if len(parts) != 2 {
			return out, fmt.Errorf("invalid line: %s", ln)
		}

		var s strategy

		if parts[0] == "A" {
			s.opponent = rock
		} else if parts[0] == "B" {
			s.opponent = paper
		} else if parts[0] == "C" {
			s.opponent = scissor
		} else {
			return out, fmt.Errorf("unable to parse opponent for line: %s", ln)
		}

		if parts[1] == "X" {
			s.me = rock
		} else if parts[1] == "Y" {
			s.me = paper
		} else if parts[1] == "Z" {
			s.me = scissor
		} else {
			return out, fmt.Errorf("unable to parse me for line: %s", ln)
		}

		out = append(out, s)
	}
	return out, nil
}

func parsePart2(in string) (guide, error) {
	var out guide
	for _, ln := range strings.Split(in, "\n") {
		ln = strings.TrimSpace(ln)

		parts := strings.Split(ln, " ")
		if len(parts) != 2 {
			return out, fmt.Errorf("invalid line: %s", ln)
		}

		var s strategy

		if parts[0] == "A" {
			s.opponent = rock
		} else if parts[0] == "B" {
			s.opponent = paper
		} else if parts[0] == "C" {
			s.opponent = scissor
		} else {
			return out, fmt.Errorf("unable to parse opponent for line: %s", ln)
		}

		if parts[1] == "X" {
			if s.opponent == rock {
				s.me = scissor
			} else if s.opponent == paper {
				s.me = rock
			} else {
				s.me = paper
			}
		} else if parts[1] == "Y" {
			s.me = s.opponent
		} else if parts[1] == "Z" {
			if s.opponent == rock {
				s.me = paper
			} else if s.opponent == paper {
				s.me = scissor
			} else {
				s.me = rock
			}
		} else {
			return out, fmt.Errorf("unable to parse me for line: %s", ln)
		}

		out = append(out, s)
	}
	return out, nil
}

type guide []strategy

func (g guide) score() int {
	var sum int
	for _, n := range g {
		sum += n.score()
	}
	return sum
}

type strategy struct {
	opponent symbol
	me       symbol
}

func (s strategy) score() int {
	if s.opponent == s.me {
		return 3 + s.me.score()
	}

	if s.opponent == rock && s.me == paper ||
		s.opponent == paper && s.me == scissor ||
		s.opponent == scissor && s.me == rock {
		return 6 + s.me.score()
	}

	return s.me.score()
}

type symbol int

const (
	rock symbol = iota
	paper
	scissor
)

func (s symbol) score() int {
	switch s {
	case rock:
		return 1
	case paper:
		return 2
	case scissor:
		return 3
	}

	return 0
}
