package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ms, err := parse(input)
	if err != nil {
		panic(err)
	}

	ms.rounds(20, 3)
	fmt.Println("monkey business after 20 rounds", ms.monkeyBusiness())

	ms2, err := parse(input)
	if err != nil {
		panic(err)
	}

	ms2.rounds(10000, 1)
	fmt.Println("monkey business after 10000 rounds", ms2.monkeyBusiness())
}

func parse(in string) (monkeys, error) {
	var ms monkeys
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
					return ms, fmt.Errorf("unable to parse %s for line %s", i, ln)
				}
				m.worries = append(m.worries, n)
			}
			continue
		}

		if strings.HasPrefix(ln, "Operation: new = ") {
			ln = strings.TrimPrefix(ln, "Operation: new = ")
			items := strings.Split(ln, " ")

			if len(items) != 3 {
				return ms, fmt.Errorf("unable to parse operation: %s", ln)
			}

			if items[0] != "old" {
				a, err := strconv.Atoi(items[0])
				if err != nil {
					return ms, fmt.Errorf("unable to parse a: %w", err)
				}
				m.a = &a
			}

			if items[2] != "old" {
				b, err := strconv.Atoi(items[2])
				if err != nil {
					return ms, fmt.Errorf("unable to parse b: %w", err)
				}
				m.b = &b
			}

			switch items[1] {
			case "*", "+":
				m.op = items[1]
			default:
				return ms, fmt.Errorf("unable to operator %s", items[1])
			}
			continue
		}

		if strings.HasPrefix(ln, "Test: divisible by ") {
			ln = strings.TrimPrefix(ln, "Test: divisible by ")
			n, err := strconv.Atoi(ln)
			if err != nil {
				return ms, fmt.Errorf("unable to parse divisble for line %s", ln)
			}

			m.mod = n
			continue
		}

		if strings.HasPrefix(ln, "If true: throw to monkey ") {
			ln = strings.TrimPrefix(ln, "If true: throw to monkey ")
			n, err := strconv.Atoi(ln)
			if err != nil {
				return ms, fmt.Errorf("unable to parse true throw to for line %s", ln)
			}

			m.yes = n
			continue
		}

		if strings.HasPrefix(ln, "If false: throw to monkey ") {
			ln = strings.TrimPrefix(ln, "If false: throw to monkey ")
			n, err := strconv.Atoi(ln)
			if err != nil {
				return ms, fmt.Errorf("unable to parse false throw to for line %s", ln)
			}

			m.no = n

			ms = append(ms, m)
			continue
		}
	}

	return ms, nil
}

type monkeys []*monkey

func (ms monkeys) monkeyBusiness() int {
	srt := ms
	sort.SliceStable(srt, func(i, j int) bool {
		return srt[i].inspected > srt[j].inspected
	})

	return srt[0].inspected * srt[1].inspected
}

func (ms monkeys) rounds(n int, relief int) {
	for i := 0; i < n; i++ {
		ms.round(relief)
	}
}

func (ms monkeys) round(relief int) {
	for _, m := range ms {
		for _, old := range m.worries {
			m.inspected++

			a := old
			if m.a != nil {
				a = *m.a
			}

			b := old
			if m.b != nil {
				b = *m.b
			}

			next := 0
			if m.op == "+" {
				next = a + b
			} else if m.op == "*" {
				next = a * b
			}

			next = (next % ms.supermod()) / relief

			to := m.no
			if next%m.mod == 0 {
				to = m.yes
			}

			ms[to].worries = append(ms[to].worries, next)
		}
		m.worries = nil
	}
}

func (ms monkeys) supermod() int {
	s := 1
	for _, m := range ms {
		s = s * m.mod
	}
	return s
}

type monkey struct {
	inspected int

	worries []int

	op string
	a  *int
	b  *int

	mod int
	yes int
	no  int
}
