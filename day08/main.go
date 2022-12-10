package main

import (
	"fmt"
	"strings"
)

func main() {
	f := parse(input)
	fmt.Println("trees visible from the edge of the forest", f.sumVisibleFromEdge())
	fmt.Println("highest scenic score", f.maxScenicScore())
}

func parse(in string) forest {
	var f forest
	for _, ln := range strings.Split(in, "\n") {
		var row []tree
		for _, c := range strings.TrimSpace(ln) {
			row = append(row, tree(c-'0'))
		}
		f = append(f, row)
	}
	return f
}

type forest [][]tree

func (f forest) sumVisibleFromEdge() int {
	var sum int
	for x, row := range f {
		for y := range row {
			if f.visibleFromEdge(x, y) {
				sum++
			}
		}
	}
	return sum
}

func (f forest) maxScenicScore() int {
	var max int
	for x, row := range f {
		for y := range row {
			s := f.scenicScore(x, y)
			if s > max {
				max = s
			}
		}
	}
	return max
}

func (f forest) visibleFromEdge(x int, y int) bool {
	notBlocked := func(_ int, blocked bool) bool {
		return !blocked
	}

	return notBlocked(f.viewDistance(up, x, y)) ||
		notBlocked(f.viewDistance(down, x, y)) ||
		notBlocked(f.viewDistance(left, x, y)) ||
		notBlocked(f.viewDistance(right, x, y))
}

func (f forest) viewDistance(d direction, x, y int) (distance int, blocked bool) {
	switch d {
	case up:
		for i := y - 1; i >= 0; i-- {
			if f.tree(x, y) <= f.tree(x, i) {
				return y - i, true
			}
		}
		return y, false

	case down:
		for i := y + 1; i < f.height(); i++ {
			if f.tree(x, y) <= f.tree(x, i) {
				return i - y, true
			}
		}
		return f.height() - 1 - y, false

	case left:
		for i := x - 1; i >= 0; i-- {
			if f.tree(x, y) <= f.tree(i, y) {
				return x - i, true
			}
		}
		return x, false

	case right:
		for i := x + 1; i < f.width(); i++ {
			if f.tree(x, y) <= f.tree(i, y) {
				return i - x, true
			}
		}
		return f.width() - 1 - x, false
	}

	return 0, true
}

func (f forest) scenicScore(x, y int) int {
	dist := func(d int, _ bool) int {
		return d
	}

	return dist(f.viewDistance(up, x, y)) *
		dist(f.viewDistance(down, x, y)) *
		dist(f.viewDistance(left, x, y)) *
		dist(f.viewDistance(right, x, y))
}

var noTree = tree(-1)

func (f forest) tree(x int, y int) tree {
	if y < 0 || y >= f.height() || x < 0 || x >= f.width() {
		return noTree
	}

	return f[y][x]
}

func (f forest) width() int {
	if f.height() == 0 {
		return 0
	}

	return len(f[0])
}

func (f forest) height() int {
	return len(f)
}

type tree int

type direction int

const (
	up direction = iota
	down
	left
	right
)

func (d direction) String() string {
	switch d {
	case up:
		return "up"
	case down:
		return "down"
	case left:
		return "left"
	case right:
		return "right"
	}

	return ""
}
