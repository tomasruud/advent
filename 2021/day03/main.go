package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	m := parse(input)

	fmt.Println(m.countTreesInPath(3, 1))

	fmt.Println(
		m.countTreesInPath(1, 1) *
			m.countTreesInPath(3, 1) *
			m.countTreesInPath(5, 1) *
			m.countTreesInPath(7, 1) *
			m.countTreesInPath(1, 2),
	)
}

func parse(in string) tobogganMap {
	return strings.Split(in, "\n")
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
