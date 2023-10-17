package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	rs, err := parse(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("sum of all intersections", rs.sumIntersections())

	sum, err := parseGroups(rs, 3).sumBadges()
	if err != nil {
		panic(err)
	}

	fmt.Println("sum of all badges", sum)
}

func parse(in string) (rucksacks, error) {
	var all rucksacks
	for _, ln := range strings.Split(in, "\n") {
		ln := strings.TrimSpace(ln)
		split := len(ln) / 2

		r := rucksack{
			compartmentA: []item(ln[:split]),
			compartmentB: []item(ln[split:]),
		}

		all = append(all, r)
	}
	return all, nil
}

func parseGroups(rs rucksacks, size int) groups {
	var gs groups
	var g group
	for i, r := range rs {
		g = append(g, r)

		if i%size == size-1 {
			gs = append(gs, g)
			g = group{}
		}
	}
	return gs
}

type groups []group

func (gs groups) sumBadges() (int, error) {
	var sum int
	for _, g := range gs {
		b, err := g.badge()
		if err != nil {
			return sum, fmt.Errorf("unable to get badge: %w", err)
		}

		sum += b.priority()
	}
	return sum, nil
}

type group rucksacks

func (g group) badge() (item, error) {
	if len(g) < 2 {
		return ' ', fmt.Errorf("unable to find badge for group: %v", g)
	}

	state := rucksack{
		compartmentA: g[0].unique(),
		compartmentB: g[1].unique(),
	}
	for _, r := range g[2:] {
		state.compartmentA = state.intersect()
		state.compartmentB = r.unique()
	}

	return state.intersect()[0], nil
}

type rucksacks []rucksack

func (g rucksacks) sumIntersections() int {
	var sum int
	for _, r := range g {
		sum += r.intersect().sum()
	}
	return sum
}

type rucksack struct {
	compartmentA items
	compartmentB items
}

func (r rucksack) unique() items {
	keys := make(map[item]bool)
	var set items
	for _, i := range r.compartmentA {
		if _, ok := keys[i]; !ok {
			keys[i] = true
			set = append(set, i)
		}
	}
	for _, i := range r.compartmentB {
		if _, ok := keys[i]; !ok {
			keys[i] = true
			set = append(set, i)
		}
	}
	return set
}

func (r rucksack) intersect() items {
	var inter items

	hash := make(map[item]bool)
	for _, e := range r.compartmentA {
		hash[e] = true
	}
	for _, e := range r.compartmentB {
		if _, ok := hash[e]; ok && hash[e] {
			inter = append(inter, e)
			hash[e] = false
		}
	}

	return inter
}

type items []item

func (is items) sum() int {
	var sum int
	for _, i := range is {
		sum += i.priority()
	}
	return sum
}

type item byte

func (i item) priority() int {
	if unicode.IsUpper(rune(i)) {
		return int(i) - int('A') + 27
	}

	return int(i) - int('a') + 1
}
