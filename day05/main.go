package main

import (
	"fmt"
	"strings"
)

func main() {
	fl := parse(input)

	fmt.Println(fl.maxSeatID())
	fmt.Println(fl.maxAvailableSeatID())
}

func parse(in string) flight {
	fl := flight{
		plane:  plane{8, 128},
		passes: []boardingPass{},
	}

	for _, l := range strings.Split(in, "\n") {
		p := boardingPass{l}
		fl.passes = append(fl.passes, p)
	}

	return fl
}

type flight struct {
	plane  plane
	passes []boardingPass
}

func (f flight) maxSeatID() int {
	var hi int
	for _, b := range f.passes {
		if f.plane.seatID(b) > hi {
			hi = f.plane.seatID(b)
		}
	}
	return hi
}

func (f flight) maxAvailableSeatID() int {
	var hi int
	for _, id := range f.availableSeatIDs() {
		if id > hi {
			hi = id
		}
	}
	return hi
}

func (f flight) availableSeatIDs() []int {
	var av []int
	for id := seatID(1, 0); id < f.maxSeatID(); id++ {
		taken := false

		for _, b := range f.passes {
			if f.plane.seatID(b) == id {
				taken = true
			}
		}

		if !taken {
			av = append(av, id)
		}
	}
	return av
}

type plane struct {
	cols int
	rows int
}

type boardingPass struct {
	seat string
}

func seatID(row int, col int) int {
	return (row * 8) + col
}

func (p plane) seatID(b boardingPass) int {
	return seatID(p.row(b), p.col(b))
}

func (p plane) row(b boardingPass) int {
	return locate(b.seat[:7], "B", 0, p.rows-1)
}

func (p plane) col(b boardingPass) int {
	return locate(b.seat[7:], "R", 0, p.cols-1)
}

func locate(code string, sig string, lo int, hi int) int {
	h := string(code[0]) == sig

	if len(code) == 1 {
		if h {
			return hi
		}

		return lo
	}

	if h {
		lo = lo + ((hi + 1 - lo) / 2)
	} else {
		hi = hi - ((hi + 1 - lo) / 2)
	}

	return locate(code[1:], sig, lo, hi)
}
