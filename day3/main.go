package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
)

var file = flag.String("file", "input.txt", "input file")

func main() {
	flag.Parse()
	var input []string

	file, err := os.Open(*file)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	m := tobogganMap{input}

	fmt.Println(m.countTreesInPath(3, 1))

	p := []int{
		m.countTreesInPath(1, 1),
		m.countTreesInPath(3, 1),
		m.countTreesInPath(5,1),
		m.countTreesInPath(7,1),
		m.countTreesInPath(1,2),
	}

	sum := 1
	for _, n := range p {
		sum *= n
	}

	fmt.Println(sum)
}

type tobogganMap struct {
	lines []string
}

func (t tobogganMap) hasTreeAt(x int, y int) (bool, error) {
	if y > t.height() {
		return false, errors.New("y coordinate is out of range")
	}

	w := len(t.lines[y])
	return string(t.lines[y][x%w]) == "#", nil
}

func (t tobogganMap) height() int {
	return len(t.lines) - 1
}

func (t tobogganMap) countTreesInPath(xStep int, yStep int) int {
	trees := 0
	for x, y := 0, 0; y < t.height(); {
		x += xStep
		y += yStep

		if yes, _ := t.hasTreeAt(x, y); yes {
			trees++
		}
	}

	return trees
}
