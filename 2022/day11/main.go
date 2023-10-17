package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	g, err := parse(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("monkey business after 20 rounds", g.withRelief(3).doRounds(20).monkeyBusiness())

	g2, err := parse(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("monkey business after 10000 rounds", g2.doRounds(10000).monkeyBusiness())
}

func parse(in string) (game, error) {
	g := game{}
	var m *monkey
	for _, ln := range strings.Split(in, "\n") {
		ln = strings.TrimSpace(ln)

		if strings.HasPrefix(ln, "Monkey ") {
			m = &monkey{}
			continue
		}

		if strings.HasPrefix(ln, "Starting items: ") {
			ln = strings.TrimPrefix(ln, "Starting items: ")
			items := strings.Split(ln, ", ")

			for _, i := range items {
				n, err := strconv.Atoi(i)
				if err != nil {
					return g, fmt.Errorf("unable to parse %s for line %s", i, ln)
				}
				m.items = append(m.items, item(n))
			}
			continue
		}

		if strings.HasPrefix(ln, "Operation: new = ") {
			ln = strings.TrimPrefix(ln, "Operation: new = ")
			items := strings.Split(ln, " ")

			if len(items) != 3 {
				return g, fmt.Errorf("unable to parse operation: %s", ln)
			}

			var a *item
			if items[0] != "old" {
				n, err := strconv.Atoi(items[0])
				if err != nil {
					return g, fmt.Errorf("unable to parse a: %w", err)
				}
				i := item(n)
				a = &i
			}

			var b *item
			if items[2] != "old" {
				n, err := strconv.Atoi(items[2])
				if err != nil {
					return g, fmt.Errorf("unable to parse b: %w", err)
				}
				i := item(n)
				b = &i
			}

			if items[1] == "*" {
				m.operation = multiply{a: a, b: b}
			} else if items[1] == "+" {
				m.operation = add{a: a, b: b}
			} else {
				return g, fmt.Errorf("unable to operator %s", items[1])
			}
			continue
		}

		if strings.HasPrefix(ln, "Test: divisible by ") {
			ln = strings.TrimPrefix(ln, "Test: divisible by ")
			n, err := strconv.Atoi(ln)
			if err != nil {
				return g, fmt.Errorf("unable to parse divisble for line %s", ln)
			}

			m.test.mod = n
			continue
		}

		if strings.HasPrefix(ln, "If true: throw to monkey ") {
			ln = strings.TrimPrefix(ln, "If true: throw to monkey ")
			n, err := strconv.Atoi(ln)
			if err != nil {
				return g, fmt.Errorf("unable to parse true throw to for line %s", ln)
			}

			m.test.yes = n
			continue
		}

		if strings.HasPrefix(ln, "If false: throw to monkey ") {
			ln = strings.TrimPrefix(ln, "If false: throw to monkey ")
			n, err := strconv.Atoi(ln)
			if err != nil {
				return g, fmt.Errorf("unable to parse false throw to for line %s", ln)
			}

			m.test.no = n

			g.monkeys = append(g.monkeys, m)
			continue
		}
	}

	return g, nil
}

type game struct {
	monkeys      []*monkey
	reliefFactor int
}

func (g game) monkeyBusiness() int {
	srt := g.monkeys
	sort.SliceStable(srt, func(i, j int) bool {
		return srt[i].inspected > srt[j].inspected
	})

	return srt[0].inspected * srt[1].inspected
}

func (g game) doRounds(n int) game {
	for i := 0; i < n; i++ {
		g.round()
	}
	return g
}

func (g game) round() {
	for _, m := range g.monkeys {
		m.inspectAll(func(i item) {
			i = g.calcRelief(i)
			to := m.throw(i)
			g.monkeys[to].catch(i)
		})
	}
}

func (g game) calcRelief(i item) item {
	mod := 1
	for _, m := range g.monkeys {
		mod = mod * m.test.mod
	}

	next := int(i) % mod
	next = next / g.relief()

	return item(next)
}

func (g game) withRelief(r int) game {
	g.reliefFactor = r
	return g
}

func (g game) relief() int {
	if g.reliefFactor < 1 {
		return 1
	}

	return g.reliefFactor
}

type monkey struct {
	inspected int
	items     []item
	operation operation
	test      test
}

func (m *monkey) inspectAll(each func(item)) {
	for range m.items {
		each(m.inspect())
	}
}

func (m *monkey) inspect() item {
	i := m.items[0]
	if len(m.items) < 2 {
		m.items = nil
	} else {
		m.items = m.items[1:]
	}

	m.inspected++

	return m.operation.calcWorry(i)
}

func (m *monkey) throw(item item) (to int) {
	return m.test.where(item)
}

func (m *monkey) catch(item item) {
	m.items = append(m.items, item)
}

type item int

type test struct {
	mod int
	yes int
	no  int
}

func (t test) where(item item) (to int) {
	if int(item)%t.mod == 0 {
		return t.yes
	}

	return t.no
}

type operation interface {
	calcWorry(old item) item
}

type add struct {
	a *item
	b *item
}

func (a add) calcWorry(old item) item {
	x := old
	if a.a != nil {
		x = *a.a
	}

	y := old
	if a.b != nil {
		y = *a.b
	}

	return x + y
}

type multiply struct {
	a *item
	b *item
}

func (m multiply) calcWorry(old item) item {
	x := old
	if m.a != nil {
		x = *m.a
	}

	y := old
	if m.b != nil {
		y = *m.b
	}

	return x * y
}
