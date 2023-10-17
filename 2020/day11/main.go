package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	g := parse(input)
	fmt.Println(g.simulateOccupied())
	fmt.Println(g.simulateOccupiedV2())
}

func parse(in string) grid {
	var g grid
	for _, l := range strings.Split(in, "\n") {
		var ps []position
		for _, p := range l {
			ps = append(ps, position(p))
		}
		g = append(g, ps)
	}
	return g
}

type position string
type grid [][]position

func (g grid) occupied() int {
	sum := 0
	for _, r := range g {
		for _, c := range r {
			if c == occupied {
				sum++
			}
		}
	}
	return sum
}

func (g grid) simulateOccupied() int {
	gr := g
	oc := gr.occupied()

	for {
		gr = gr.simulate()
		o := gr.occupied()

		if o == oc {
			return oc
		}

		oc = o
	}
}

func (g grid) copy() grid {
	cp := make(grid, len(g))
	for i := range g {
		cp[i] = make([]position, len(g[i]))
		copy(cp[i], g[i])
	}
	return cp
}

func (g grid) simulate() grid {
	cp := g.copy()
	for y := range g {
		for x := range g[y] {
			if g[y][x] == empty && g.adjacent(x, y).occupied() == 0 {
				cp[y][x] = occupied
			} else if g[y][x] == occupied && g.adjacent(x, y).occupied() >= 4 {
				cp[y][x] = empty
			}
		}
	}
	return cp
}

func (g grid) adjacent(x int, y int) grid {
	var pos []position
	for i := 0; i < 9; i++ {
		if i == 4 {
			// i is x,y
			continue
		}

		x1 := x - 1 + (i % 3)
		y1 := y - 1 + (i / 3)

		p, err := g.get(x1, y1)

		if err != nil {
			// assume out of bounds
			continue
		}

		pos = append(pos, p)
	}

	return grid{pos}
}

func (g grid) closest(x int, y int) grid {
	pos := make(map[int]position)

	for z := 1; ; z++ {
		for i := 0; i < 9; i++ {
			if _, exist := pos[i]; exist || i == 4 {
				// i is x,y
				continue
			}

			x1 := x - z + (i%3)*z
			y1 := y - z + (i/3)*z

			p, err := g.get(x1, y1)

			if err != nil {
				pos[i] = floor
				continue
			}

			if p != floor {
				pos[i] = p
			}
		}

		if len(pos) >= 8 {
			break
		}
	}

	var out []position
	for _, v := range pos {
		out = append(out, v)
	}
	return grid{out}
}

func (g grid) get(x int, y int) (position, error) {
	if x < 0 || y < 0 || x >= len(g[0]) || y >= len(g) {
		// out of bounds
		return "", errors.New("out of bounds")
	}

	return g[y][x], nil
}

func (g grid) simulateV2() grid {
	cp := g.copy()
	for y := range g {
		for x := range g[y] {
			if g[y][x] == empty && g.closest(x, y).occupied() == 0 {
				cp[y][x] = occupied
			} else if g[y][x] == occupied && g.closest(x, y).occupied() >= 5 {
				cp[y][x] = empty
			}
		}
	}
	return cp
}

func (g grid) simulateOccupiedV2() int {
	gr := g
	oc := gr.occupied()

	for {
		gr = gr.simulateV2()
		o := gr.occupied()

		if o == oc {
			return oc
		}

		oc = o
	}
}