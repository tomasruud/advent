package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b := parse(input)
	fmt.Println(b.diffSize())
	fmt.Println(b.branches())
}

func parse(in string) bag {
	b := bag{}
	for _, v := range strings.Split(in, "\n") {
		i, _ := strconv.Atoi(v)
		b = append(b, adapter(i))
	}
	return b
}

type bag []adapter

func (b bag) branches() int {
	cache := make(map[adapter]int)
	return b.chain().cachedBranches(cache)
}

func (b bag) cachedBranches(cache map[adapter]int) int {
	if len(b) < 3 {
		return 1
	}

	z := b[0]
	if le, exist := cache[z]; exist {
		return le
	}

	cache[z] = 0
	for i, v := range b[1:] {
		if z.connects(v) {
			cache[z] += b[1+i:].cachedBranches(cache)
		} else {
			break
		}
	}

	return cache[z]
}

func (b bag) sort() bag {
	cp := make(bag, len(b))
	copy(cp, b)

	sort.Slice(cp, func(i, j int) bool {
		return cp[i] < cp[j]
	})

	return cp
}

func (b bag) chain() bag {
	c := bag{0}
	for _, j := range b.sort() {
		if !c[len(c)-1].connects(j) {
			break
		}

		c = append(c, j)
	}

	return append(c, c[len(c)-1]+3)
}

func (b bag) diffSize() int {
	c := b.chain()

	n1, n3 := 0, 0
	for i, j := range c[1:] {
		d := c[i].diff(j)

		if d == 1 {
			n1++
		} else if d == 3 {
			n3++
		}
	}

	return n1 * n3
}

type adapter int

func (a adapter) diff(o adapter) int {
	return int(o) - int(a)
}

func (a adapter) connects(o adapter) bool {
	return 0 <= a.diff(o) && a.diff(o) <= 3
}
