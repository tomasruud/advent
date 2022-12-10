package main

import (
	"fmt"
	"strings"
)

func main() {
	f := parse(input)

	fmt.Println(f.uniqueSum())
	fmt.Println(f.allSum())
}

func parse(in string) flight {
	f := flight{}
	g := group{}
	for _, line := range strings.Split(in, "\n") {
		if line == "" {
			f = append(f, g)
			g = group{}
			continue
		}

		fo := form{}
		for _, yes := range line {
			fo = append(fo, string(yes))
		}
		g = append(g, fo)
	}

	f = append(f, g)

	return f
}

type flight []group

func (f flight) uniqueSum() int {
	tot := 0
	for _, g := range f {
		tot += g.unique()
	}
	return tot
}

func (f flight) allSum() int {
	tot := 0
	for _, g := range f {
		tot += g.all()
	}
	return tot
}

type group []form
type form []string

func (g group) unique() int {
	set := make(map[string]bool)
	for _, f := range g {
		for _, a := range f {
			set[a] = true
		}
	}
	return len(set)
}

func (g group) all() int {
	set := make(map[string]int)
	for _, f := range g {
		for _, a := range f {
			set[a] = set[a] + 1
		}
	}

	tot := 0
	for _, v := range set {
		if v == len(g) {
			tot++
		}
	}
	return tot
}
