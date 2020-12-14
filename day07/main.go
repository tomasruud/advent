package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	bs := parse(input)
	fmt.Println(bs.contains("shiny gold"))
	fmt.Println(bs.find("shiny gold").size())
}

func parse(in []string) bags {
	bs := make(map[kind]*bag)
	roots := make(map[kind]*bag)

	rf := regexp.MustCompile(`^([a-z]+ [a-z]+) bags contain`)
	rs := regexp.MustCompile(`(\d+) ([a-z]+ [a-z]+) bags?`)

	for _, line := range in {
		k := kind(rf.FindStringSubmatch(line)[1])
		if _, exist := bs[k]; !exist {
			bs[k] = newBag(k)
		}

		roots[k] = bs[k]

		for _, ks := range rs.FindAllStringSubmatch(line, -1) {
			count, _ := strconv.Atoi(ks[1])
			k2 := kind(ks[2])

			if _, exist := bs[k2]; !exist {
				bs[k2] = newBag(k2)
			}

			delete(roots, k2)

			for i := 0; i < count; i++ {
				bs[k].bags = append(bs[k].bags, bs[k2])
			}
		}
	}

	out := newBags()
	for _, r := range roots {
		out = append(out, r)
	}
	return out
}

type kind string

type bags []*bag

func newBags(bs ...*bag) bags {
	return bs
}

func (bs bags) all() bags {
	if len(bs) < 1 {
		return newBags()
	}

	ex := make(map[kind]bool)
	types := newBags()
	for _, b := range bs {
		if _, exist := ex[b.kind]; exist {
			continue
		}

		ex[b.kind] = true
		types = append(types, b)

		for _, n := range b.bags.all() {
			if _, exist := ex[n.kind]; exist {
				continue
			}

			ex[n.kind] = true
			types = append(types, n)
		}
	}
	return types
}

func (bs bags) contains(k kind) int {
	sum := 0
	all := bs.all()

	for _, b := range all {
		if b.bags.find(k) != nil {
			sum++
		}
	}
	return sum
}

func (bs bags) find(k kind) *bag {
	if len(bs) < 1 {
		return nil
	}

	visited := make(map[kind]bool)

	for _, b := range bs {
		if _, exist := visited[b.kind]; exist {
			continue
		}

		if b.kind == k {
			return b
		}

		if bb := b.bags.find(k); bb != nil {
			return bb
		}

		visited[b.kind] = true
	}

	return nil
}

type bag struct {
	kind kind
	bags bags
}

func newBag(k kind, bs ...*bag) *bag {
	return &bag{
		kind: k,
		bags: bs,
	}
}

func (b *bag) size() int {
	sum := len(b.bags)
	for _, b := range b.bags {
		sum += b.size()
	}
	return sum
}
