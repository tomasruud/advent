package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("provide input file as argument")
	}

	loc := args[0]

	nums, err := read2(loc)

	if err != nil {
		log.Fatalf("unable to read input %v", err)
	}

	res := find(nums, 2020)

	if len(res) != 3 {
		log.Fatalf("result lenght too short %v", res)
	}

	a, b, c := res[0], res[1], res[2]

	if a+b+c != 2020 {
		log.Fatalf("a+b+c was not 2020")
	}

	log.Printf("result of %d + %d + %d = %d\n", a, b, c, a+b+c)
	log.Printf("result of %d * %d * %d = %d", a, b, c, a*b*c)
}

func find(nums []int, target int) []int {
	for i, a := range nums {
		anums := remove(nums, i)

		for j, b := range anums {
			bnums := remove(anums, j)

			for _, c := range bnums {
				if a+b+c == target {
					return []int{a, b, c}
				}
			}
		}
	}

	return []int{}
}

func remove(s []int, i int) []int {
	c := make([]int, len(s))
	copy(c, s)

	c[i] = c[len(c)-1]
	return c[:len(c)-1]
}

func read2(loc string) ([]int, error) {
	var nums []int

	file, err := os.Open(loc)

	if err != nil {
		return nums, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return nums, err
		}

		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		return nums, err
	}

	return nums, nil
}
