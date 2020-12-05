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
	fl := flight{
		plane:  plane{8, 128},
		passes: []boardingPass{},
	}

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
		p := boardingPass{scanner.Text(), fl.plane}
		fl.passes = append(fl.passes, p)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(fl.maxSeatID())
	fmt.Println(fl.maxAvailableSeatID())
}

type flight struct {
	plane  plane
	passes []boardingPass
}

func (f flight) maxSeatID() int {
	var hi int
	for _, b := range f.passes {
		if b.seatID() > hi {
			hi = b.seatID()
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
			if b.seatID() == id {
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
	seatCode string
	plane    plane
}

func seatID(row int, col int) int {
	return (row * 8) + col
}

func (b boardingPass) seatID() int {
	return seatID(b.row(), b.col())
}

func (b boardingPass) row() int {
	return locate(b.seatCode[:7], "B", 0, b.plane.rows-1)
}

func (b boardingPass) col() int {
	return locate(b.seatCode[7:], "R", 0, b.plane.cols-1)
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
