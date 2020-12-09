package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

	x := parse(input, 25)
	fmt.Println(x.invalid())
	fmt.Println(x.weak())
}

func parse(in []string, pre int) xmas {
	x := xmas{pre: pre}

	for _, i := range in {
		v, _ := strconv.Atoi(i)
		x.list = append(x.list, v)
	}

	return x
}

type xmas struct {
	pre  int
	list []int
}

func (x xmas) valid(v int, i int) bool {
	for p1, v1 := range x.list[i-x.pre : i] {
		for p2, v2 := range x.list[i-x.pre : i] {
			if p1 == p2 {
				continue
			}

			if v1+v2 == v {
				return true
			}
		}
	}

	return false
}

func (x xmas) invalid() int {
	for i, n := range x.list[x.pre:] {
		if !x.valid(n, x.pre+i) {
			return n
		}
	}

	return -1
}

func (x xmas) weak() int {
	var sta int
	var end int
	var v1 int
	var v2 int
	inv := x.invalid()

	for sta, v1 = range x.list {
		if inv == v1 {
			continue
		}

		sum := v1
		for end, v2 = range x.list[sta+1:] {
			if inv == v2 {
				break
			}

			sum += v2

			if sum == inv || sum > inv {
				break
			}
		}

		if sum == inv {
			break
		}
	}

	min, max := x.list[sta], x.list[sta]

	for _, v := range x.list[sta : sta+end+1] {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return max + min
}
