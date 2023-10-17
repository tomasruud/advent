package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	s, ms, err := parse(input)
	if err != nil {
		panic(err)
	}

	if err = s.apply(ms, crateMover9000); err != nil {
		panic(err)
	}

	fmt.Println("top of all stacks with CrateMover9000", s.top())

	s1, ms, err := parse(input)
	if err != nil {
		panic(err)
	}

	if err := s1.apply(ms, crateMover9001); err != nil {
		panic(err)
	}

	fmt.Println("top of all stacks with CrateMover9001", s1.top())
}

func parse(in string) (ship, moves, error) {
	var s ship
	var ms moves

	stackRegexp := regexp.MustCompile(`(?: {3}|\[([A-Z]+)]) ?`)
	moveRegexp := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

	step := "crates"
	for _, ln := range strings.Split(in, "\n") {
		if !strings.Contains(ln, "[") && step == "crates" {
			step = "ids"
			continue
		} else if strings.TrimSpace(ln) == "" {
			step = "moves"
			continue
		}

		switch step {
		case "crates":
			for i, matches := range stackRegexp.FindAllStringSubmatch(ln, -1) {
				if i >= len(s.stacks) {
					s.stacks = append(s.stacks, nil)
				}

				if len(matches) < 2 {
					return s, ms, fmt.Errorf("unable to parse line: %s", ln)
				}

				if matches[1] != "" {
					s.stacks[i] = append(s.stacks[i], crate(matches[1]))
				}
			}

		case "moves":
			var m move
			for i, match := range moveRegexp.FindStringSubmatch(ln) {
				if i == 0 {
					continue
				}

				n, err := strconv.Atoi(match)
				if err != nil {
					return s, ms, fmt.Errorf("unable to parse move value: %s", match)
				}

				if i == 1 {
					m.amount = n
				} else if i == 2 {
					m.from = n
				} else if i == 3 {
					m.to = n
				}
			}
			ms = append(ms, m)
		}
	}

	return s, ms, nil
}

type moves []move

type move struct {
	amount int
	from   int
	to     int
}

func (m move) fromIndex() int {
	return m.from - 1
}

func (m move) toIndex() int {
	return m.to - 1
}

type ship struct {
	stacks []crates
}

func (s ship) top() crates {
	var top crates
	for _, stack := range s.stacks {
		top = append(top, stack.top())
	}
	return top
}

type crateMover func(*ship, moves) error

func crateMover9000(s *ship, ms moves) error {
	for _, m := range ms {
		if m.fromIndex() >= len(s.stacks) {
			return fmt.Errorf("trying to move from non existing stack %d", m.from)
		}

		if m.toIndex() >= len(s.stacks) {
			return fmt.Errorf("trying to move to non existing stack %d", m.to)
		}

		if len(s.stacks[m.fromIndex()]) < m.amount {
			return fmt.Errorf("%d does not have enough elements to move %d", m.from, m.amount)
		}

		for i := 0; i < m.amount; i++ {
			el := s.stacks[m.fromIndex()][0]
			s.stacks[m.fromIndex()] = s.stacks[m.fromIndex()][1:]
			s.stacks[m.toIndex()] = append(crates{el}, s.stacks[m.toIndex()]...)
		}
	}
	return nil
}

func crateMover9001(s *ship, ms moves) error {
	for _, m := range ms {
		if m.fromIndex() >= len(s.stacks) {
			return fmt.Errorf("trying to move from non existing stack %d", m.from)
		}

		if m.toIndex() >= len(s.stacks) {
			return fmt.Errorf("trying to move to non existing stack %d", m.to)
		}

		if len(s.stacks[m.fromIndex()]) < m.amount {
			return fmt.Errorf("%d does not have enough elements to move %d", m.from, m.amount)
		}

		els := s.stacks[m.fromIndex()][0:m.amount]
		s.stacks[m.fromIndex()] = s.stacks[m.fromIndex()][m.amount:]

		var next crates
		next = append(next, els...)
		s.stacks[m.toIndex()] = append(next, s.stacks[m.toIndex()]...)
	}
	return nil
}

func (s ship) apply(ms moves, st crateMover) error {
	if err := st(&s, ms); err != nil {
		return fmt.Errorf("unable to apply moves: %w", err)
	}

	return nil
}

type crates []crate

func (cs crates) String() string {
	var s string
	for _, c := range cs {
		s += string(c)
	}
	return s
}

func (cs crates) top() crate {
	return cs[0]
}

type crate string
