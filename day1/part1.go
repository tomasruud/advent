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

	nums, err := read(loc)

	if err != nil {
		log.Fatalf("unable to read input %v", err)
	}

	var a int
	var b int

	for range nums {
		a, nums = nums[0], nums[1:]

		for _, num := range nums {
			if (num + a) == 2020 {
				b = num
				break
			}
		}

		if b > 0 {
			break
		}
	}

	if a+b != 2020 {
		log.Fatalf("a+b was not 2020")
	}

	log.Printf("result of %d * %d = %d", a, b, a*b)
}

func read(loc string) ([]int, error) {
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
