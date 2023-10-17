package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ms, err := parse(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("total unique tail positions")
	fmt.Println("for rope with length 2 ", len(newRope(2).move(ms)))
	fmt.Println("for rope with length 10", len(newRope(10).move(ms)))
}

func parse(in string) ([]move, error) {
	var ms []move
	for _, ln := range strings.Split(in, "\n") {
		ln = strings.TrimSpace(ln)

		pt := strings.Split(ln, " ")
		if len(pt) != 2 {
			return ms, fmt.Errorf("unable to parse line: %s", ln)
		}

		n, err := strconv.Atoi(pt[1])
		if err != nil {
			return ms, fmt.Errorf("unable to parse length for line: %s", ln)
		}

		var m move
		switch pt[0] {
		case "D":
			m = down
		case "U":
			m = up
		case "L":
			m = left
		case "R":
			m = right
		}

		for i := 0; i < n; i++ {
			ms = append(ms, m)
		}
	}
	return ms, nil
}

func newRope(n int) rope {
	var r rope
	for i := 0; i < n; i++ {
		r = append(r, &knot{x: 0, y: 0})
	}
	return r
}

type rope []*knot

func (r rope) move(ms []move) (trail []knot) {
	u := make(map[knot]bool)
	for _, m := range ms {
		r[0].move(m)
		r[1:].follow(r[0])

		end := *r[len(r)-1]
		if _, ok := u[end]; !ok {
			u[end] = true
			trail = append(trail, end)
		}
	}
	return trail
}

func (r rope) follow(head *knot) {
	if len(r) == 0 || r[0].adjacent(head) {
		return
	}

	if head.x < r[0].x {
		r[0].move(left)
	} else if head.x > r[0].x {
		r[0].move(right)
	}

	if head.y < r[0].y {
		r[0].move(down)
	} else if head.y > r[0].y {
		r[0].move(up)
	}

	if len(r) > 1 {
		r[1:].follow(r[0])
	}
}

type move int

const (
	up move = iota
	down
	left
	right
)

type knot struct {
	x int
	y int
}

func (k *knot) move(m move) {
	switch m {
	case up:
		k.y++
	case down:
		k.y--
	case left:
		k.x--
	case right:
		k.x++
	}
}

func (k *knot) adjacent(other *knot) bool {
	if k.x != other.x &&
		k.x != other.x-1 &&
		k.x != other.x+1 {
		return false
	}

	if k.y != other.y &&
		k.y != other.y+1 &&
		k.y != other.y-1 {
		return false
	}

	return true
}
