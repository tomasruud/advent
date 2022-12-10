package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	ts, sc := parse(input)
	fmt.Println(sc.next(ts))
	fmt.Println(parseV2(input).firstChain(100000000000000))
}

func parse(in string) (timestamp, schedule) {
	i := strings.Split(in, "\n")

	ts, _ := strconv.Atoi(i[0])

	sc := schedule{}
	for _, b := range strings.Split(i[1], ",") {
		if b == "x" {
			continue
		}

		bu, _ := strconv.Atoi(b)
		sc = append(sc, bus(bu))
	}

	return timestamp(ts), sc
}

const any = -1

func parseV2(in string) schedule {
	i := strings.Split(in, "\n")

	sc := schedule{}
	for _, b := range strings.Split(i[1], ",") {
		if b == "x" {
			sc = append(sc, any)
			continue
		}

		bu, _ := strconv.Atoi(b)
		sc = append(sc, bus(bu))
	}

	return sc
}

type schedule []bus
type timestamp int
type bus int

func (b bus) departsAt(t timestamp) bool {
	return int(t)%int(b) == 0
}

func (s schedule) nextBus(a timestamp) (bus, timestamp) {
	for n := a; ; n++ {
		for _, b := range s {
			if b.departsAt(n) {
				return b, n
			}
		}
	}
}

func (s schedule) next(a timestamp) int {
	b, t := s.nextBus(a)
	return int(t-a) * int(b)
}

func (s schedule) chains(t timestamp) bool {
	for i, b := range s {
		if b == any {
			continue
		}

		if !b.departsAt(t + timestamp(i)) {
			return false
		}
	}
	return true
}

type tup map[int]bus

func (ts tup) chains(t timestamp) bool {
	for i, b := range ts {
		if !b.departsAt(t + timestamp(i)) {
			return false
		}
	}
	return true
}

const concurrency = 10
func (s schedule) firstChain(t timestamp) timestamp {
	max, mi := s[0], uint64(0)
	ts := make(tup)

	for i, m := range s {
		if m != any {
			ts[i] = m
		}

		if m > max {
			max = m
			mi = uint64(i)
		}
	}

	n := uint64(t) / uint64(max)

	for {
		var wg sync.WaitGroup

		for x := 0; x < concurrency; x++ {
			if ts.chains(timestamp(n * uint64(max) - mi)) {
				return timestamp(n * uint64(max) - mi)
			}

			n++
		}

		wg.Wait()
	}
}
