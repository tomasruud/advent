package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	es, err := parse(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("most calories", es.topCalories(1).totalCalories())
	fmt.Println("top 3 calories", es.topCalories(3).totalCalories())
}

func parse(input string) (elves, error) {
	var all elves
	var current elf
	for _, ln := range strings.Split(input, "\n") {
		ln = strings.TrimSpace(ln)

		if ln == "" {
			all = append(all, current)
			current = elf{}
			continue
		}

		n, err := strconv.Atoi(ln)
		if err != nil {
			return all, fmt.Errorf("unable to parse number: %w", err)
		}

		current.calories = append(current.calories, n)
	}

	return all, nil
}

type elves []elf

func (es elves) topCalories(n int) elves {
	sort.Slice(es, func(i, j int) bool {
		return es[i].totalCalories() > es[j].totalCalories()
	})

	return es[:n]
}

func (es elves) totalCalories() int {
	var sum int
	for _, e := range es {
		sum += e.totalCalories()
	}
	return sum
}

type elf struct {
	calories []int
}

func (e elf) totalCalories() int {
	var sum int
	for _, cal := range e.calories {
		sum += cal
	}
	return sum
}
