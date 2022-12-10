package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ps, err := parse(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("total fully contained", ps.sumFullyContains())
	fmt.Println("total overlaps", ps.sumOverlaps())
}

func parse(in string) (pairs, error) {
	var ps pairs
	for _, ln := range strings.Split(in, "\n") {
		ln = strings.TrimSpace(ln)

		var p pair
		for _, rs := range strings.Split(ln, ",") {
			r := strings.Split(rs, "-")
			if len(r) != 2 {
				return ps, fmt.Errorf("unable to parse range: %s", r)
			}

			start, err := strconv.Atoi(r[0])
			if err != nil {
				return ps, fmt.Errorf("unable to parse range start: %w", err)
			}

			end, err := strconv.Atoi(r[1])
			if err != nil {
				return ps, fmt.Errorf("unable to parse range stop: %w", err)
			}

			var e elf
			for i := start; i <= end; i++ {
				e.sections = append(e.sections, section(i))
			}
			p = append(p, e)
		}
		ps = append(ps, p)
	}
	return ps, nil
}

type pairs []pair

func (ps pairs) sumFullyContains() int {
	var sum int
	for _, p := range ps {
		if p.fullyContains() {
			sum++
		}
	}
	return sum
}

func (ps pairs) sumOverlaps() int {
	var sum int
	for _, p := range ps {
		if p.overlaps() {
			sum++
		}
	}
	return sum
}

type pair []elf

func (p pair) fullyContains() bool {
	if len(p) < 2 {
		return true
	}

	sort.SliceStable(p, func(i, j int) bool {
		return len(p[i].sections) > len(p[j].sections)
	})

	a := p[0]
	for _, b := range p[1:] {
		if !a.hasSection(b.sections) {
			return false
		}
		a = b
	}

	return true
}

func (p pair) overlaps() bool {
	if len(p) < 2 {
		return true
	}

	a := p[0]
	for _, b := range p[1:] {
		if !a.overlaps(b) {
			return false
		}
		a = b
	}

	return true
}

type elf struct {
	sections []section
}

func (e elf) overlaps(eo elf) bool {
	for _, s := range eo.sections {
		if e.contains(s) {
			return true
		}
	}
	return false
}

func (e elf) contains(s section) bool {
	for _, a := range e.sections {
		if a == s {
			return true
		}
	}
	return false
}

func (e elf) hasSection(ss []section) bool {
	if len(e.sections) < len(ss) {
		return false
	}

	for _, s := range ss {
		if !e.contains(s) {
			return false
		}
	}
	return true
}

type section int
