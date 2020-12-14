package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var file = flag.String("file", "input.txt", "input file")

func main() {
	flag.Parse()
	var input []int

	file, err := os.Open(*file)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic(err)
		}

		input = append(input, num)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	d := day1{
		input:  input,
		target: 2020,
	}

	fmt.Println(d.task1())
	fmt.Println(d.task2())
}

type day1 struct {
	input  []int
	target int
}

func (d day1) task1() (int, error) {
	for i, a := range d.input {
		anums := remove(d.input, i)
		for _, b := range anums {
			if a+b == d.target {
				return a * b, nil
			}
		}
	}

	return -1, errors.New("no matching numbers found")
}

func (d day1) task2() (int, error) {
	for i, a := range d.input {
		anums := remove(d.input, i)

		for j, b := range anums {
			bnums := remove(anums, j)

			for _, c := range bnums {
				if a+b+c == d.target {
					return a * b * c, nil
				}
			}
		}
	}

	return -1, errors.New("no matching numbers found")
}

func remove(s []int, i int) []int {
	c := make([]int, len(s))
	copy(c, s)

	c[i] = c[len(c)-1]
	return c[:len(c)-1]
}
