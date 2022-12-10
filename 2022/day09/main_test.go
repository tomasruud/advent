package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fixture = []move{
	right, right, right, right,
	up, up, up, up,
	left, left, left,
	down,
	right, right, right, right,
	down,
	left, left, left, left, left,
	right, right,
}

func Test_parse(t *testing.T) {
	t.Run("it parses test input", func(t *testing.T) {
		got, err := parse(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`)

		assert.NoError(t, err)
		assert.Equal(t, fixture, got)
	})
}

func Test_knot_adjacent(t *testing.T) {
	tests := []struct {
		name  string
		p     *knot
		other *knot
		want  bool
	}{
		{p: &knot{0, 0}, other: &knot{0, 1}, want: true, name: "adjacent"},
		{p: &knot{0, 0}, other: &knot{1, 0}, want: true, name: "adjacent"},
		{p: &knot{0, 0}, other: &knot{0, -1}, want: true, name: "adjacent"},
		{p: &knot{0, 0}, other: &knot{-1, 0}, want: true, name: "adjacent"},
		{p: &knot{0, 0}, other: &knot{-1, -1}, want: true, name: "adjacent"},
		{p: &knot{0, 0}, other: &knot{-1, 1}, want: true, name: "adjacent"},
		{p: &knot{0, 0}, other: &knot{-2, 0}, want: false, name: "not adjacent"},
		{p: &knot{0, 0}, other: &knot{0, -2}, want: false, name: "not adjacent"},
		{p: &knot{0, 0}, other: &knot{0, 2}, want: false, name: "not adjacent"},
		{p: &knot{0, 0}, other: &knot{2, 0}, want: false, name: "not adjacent"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.p.adjacent(tt.other))
		})
	}
}

func Test_knot_apply(t *testing.T) {
	tests := []struct {
		name string
		k    knot
		m    move
		want knot
	}{
		{name: "applies", k: knot{}, m: right, want: knot{x: 1}},
		{name: "applies", k: knot{}, m: left, want: knot{x: -1}},
		{name: "applies", k: knot{}, m: up, want: knot{y: 1}},
		{name: "applies", k: knot{}, m: down, want: knot{y: -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.move(tt.m)
			assert.Equal(t, tt.want, tt.k)
		})
	}
}

func Test_rope_move(t *testing.T) {
	t.Run("rope with length 2", func(t *testing.T) {
		assert.Equal(t, 13, len(newRope(2).move(fixture)))
	})

	t.Run("rope with length 10", func(t *testing.T) {
		assert.Equal(t, 1, len(newRope(10).move(fixture)))
	})

	t.Run("rope with length 10 but new moves", func(t *testing.T) {
		f, _ := parse(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`)
		assert.Equal(t, 36, len(newRope(10).move(f)))
	})
}
