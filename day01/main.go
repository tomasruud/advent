package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	r := parse(input)

	fmt.Println(r.double(2020))
	fmt.Println(r.triple(2020))
}

func parse(in string) report {
	var r report
	for _, l := range strings.Split(in, "\n") {
		a, err := strconv.Atoi(l)

		if err != nil {
			continue
		}

		r = append(r, a)
	}
	return r
}

type report []int

func (r report) double(v int) (int, error) {
	for i := 0; i < len(r); i++ {
		a := r[i]
		for _, b := range r[i+1:] {
			if a+b == v {
				return a * b, nil
			}
		}
	}

	return -1, errors.New("no integers found")
}

func (r report) triple(v int) (int, error) {
	for i := 0; i < len(r); i++ {
		a := r[i]
		for j := i + 1; j < len(r); j++ {
			b := r[j]
			for _, c := range r[j+1:] {
				if a+b+c == v {
					return a * b * c, nil
				}
			}
		}
	}

	return -1, errors.New("no integers found")
}
