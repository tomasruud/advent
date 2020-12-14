package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var file = flag.String("file", "input.txt", "input file")

func main() {
	flag.Parse()
	var input []string

	file, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	f := parse(input)
	fmt.Println(f.uniqueSum())
	fmt.Println(f.allSum())
}

func parse(i []string) flight {
	f := flight{}
	g := group{}
	for _, line := range i {
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
