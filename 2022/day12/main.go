package main

import "strings"

func main() {
}

func parse(in string) heightmap {
	var h heightmap
	for i, ln := range strings.Split(in, "\n") {
		ln = strings.TrimSpace(ln)
		if len(h) >= i {
			h = append(h, []rune(ln))
		}
	}
	return h
}

type heightmap [][]rune

func (h heightmap) paths() [][]bool {
	var p [][]bool
	for y := range h {
		var row []bool
		for x := range h[y] {
			height := h.height(x, y)

		}
	}
	return p
}

func (h heightmap) start() (x, y int) {
	for y := range h {
		for x, el := range h[y] {
			if el == 'S' {
				return x, y
			}
		}
	}
	return -1, -1
}

func (h heightmap) end() (x, y int) {
	for y := range h {
		for x, el := range h[y] {
			if el == 'E' {
				return x, y
			}
		}
	}
	return -1, -1
}

func (h heightmap) height(x, y int) int {
	if y < 0 || y >= len(h) {
		return -1
	}

	if x < 0 || x >= len(h) {
		return -1
	}

	if h[y][x] == 'S' || h[y][x] == 'E' {
		return 0
	}

	return int(h[y][x]-'a') + 1
}

type tree struct {
	root node
}

type node struct {
}
