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
	m := tobogganMap{}

	file, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	defer func () {
		err := file.Close()
		panic(err)
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m = append(m, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(m.countTreesInPath(3, 1))

	fmt.Println(
		m.countTreesInPath(1, 1) *
			m.countTreesInPath(3, 1) *
			m.countTreesInPath(5, 1) *
			m.countTreesInPath(7, 1) *
			m.countTreesInPath(1, 2),
	)
}

type tobogganMap []string

func (t tobogganMap) hasTreeAt(x int, y int) (bool, error) {
	if y > t.height() {
		return false, errors.New("y coordinate is out of range")
	}

	w := len(t[y])
	return string(t[y][x%w]) == "#", nil
}

func (t tobogganMap) height() int {
	return len(t) - 1
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
