package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	ms := parse(input)
	s := ship{0, 0, east}
	fmt.Println(s.moves(ms).manhattan())

	w := waypoint{10, 1, s}
	fmt.Println(w.moves(ms).s.manhattan())
}

func parse(in string) []move {
	r := regexp.MustCompile(`^([NSEWLRF])(\d+)$`)

	var m []move
	for _, l := range strings.Split(in, "\n") {
		c := r.FindStringSubmatch(l)

		if len(c) != 3 {
			continue
		}

		v, _ := strconv.Atoi(c[2])
		m = append(m, move{c[1], v})
	}
	return m
}

type move struct {
	direction string
	value     int
}

type ship struct {
	x         int
	y         int
	direction string
}

func (s ship) manhattan() int {
	return abs(s.x) + abs(s.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func (s ship) moves(ms []move) ship {
	n := s
	for _, m := range ms {
		n = n.move(m)
	}
	return n
}

func (s ship) move(m move) ship {
	n := ship{
		s.x,
		s.y,
		s.direction,
	}

	if m.direction == right || m.direction == left {
		for i := 0; i < (m.value / 90); i++ {
			switch n.direction {
			case north:
				if m.direction == right {
					n.direction = east
				} else {
					n.direction = west
				}
			case east:
				if m.direction == right {
					n.direction = south
				} else {
					n.direction = north
				}
			case south:
				if m.direction == right {
					n.direction = west
				} else {
					n.direction = east
				}
			case west:
				if m.direction == right {
					n.direction = north
				} else {
					n.direction = south
				}
			}
		}
		return n
	}

	dir := m.direction

	if dir == forward {
		dir = n.direction
	}

	switch dir {
	case north:
		n.y += m.value
	case east:
		n.x += m.value
	case south:
		n.y -= m.value
	case west:
		n.x -= m.value
	}

	return n
}

type waypoint struct {
	x int
	y int

	s ship
}

func (w waypoint) move(m move) waypoint {
	n := waypoint{
		w.x,
		w.y,
		w.s,
	}

	if m.direction == forward {
		n.s.x += w.x * m.value
		n.s.y += w.y * m.value
		return n
	}

	if m.direction == right {
		m.direction = left
		m.value = 360 - (((m.value / 90) % 4) * 90)
	}

	if m.direction == left {
		for i := 0; i < (m.value/90)%4; i++ {
			tmp := n.x
			n.x = -n.y
			n.y = tmp
		}
		return n
	}

	switch m.direction {
	case north:
		n.y += m.value
	case south:
		n.y -= m.value
	case west:
		n.x -= m.value
	case east:
		n.x += m.value
	}

	return n
}

func (w waypoint) moves(ms []move) waypoint {
	n := w
	for _, m := range ms {
		n = n.move(m)
	}
	return n
}
