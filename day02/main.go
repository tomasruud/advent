package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var file = flag.String("file", "input.txt", "input file")

func main() {
	flag.Parse()
	var input []pw

	file, err := os.Open(*file)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		p, err := parseEntry(scanner.Text())

		if err != nil {
			panic(err)
		}

		input = append(input, p)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	d := day2{input}

	fmt.Println(d.task1())
	fmt.Println(d.task2())
}

type pw struct {
	min  int
	max  int
	char string
	pass string
}

func parseEntry(e string) (pw, error) {
	r := regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<char>[A-Za-z]): (?P<pass>[A-Za-z]+)`)

	match := r.FindStringSubmatch(e)
	sub := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" && i < len(match) {
			sub[name] = match[i]
		}
	}

	min, err := strconv.Atoi(sub["min"])

	if err != nil {
		return pw{}, err
	}

	max, err := strconv.Atoi(sub["max"])

	if err != nil {
		return pw{}, err
	}

	return pw{
		min:  min,
		max:  max,
		char: sub["char"],
		pass: sub["pass"],
	}, nil
}

func (p pw) valid() bool {
	c := strings.Count(p.pass, p.char)

	if c >= p.min && c <= p.max {
		return true
	}

	return false
}

func (p pw) validAtToboggan() bool {
	l := len(p.pass)
	first := p.min - 1
	second := p.max - 1

	matched := false

	if first >= 0 && first < l && string(p.pass[first]) == p.char {
		matched = true
	}

	if second >= 0 && second < l && string(p.pass[second]) == p.char {
		matched = !matched
	}

	return matched
}

type day2 struct {
	input []pw
}

// task1 returns number of passwords valid
func (d day2) task1() int {
	tot := 0
	for _, p := range d.input {
		if p.valid() {
			tot++
		}
	}

	return tot
}

// task2 returns number of passwords valid at Toboggan
func (d day2) task2() int {
	tot := 0
	for _, p := range d.input {
		if p.validAtToboggan() {
			tot++
		}
	}

	return tot
}
