package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const inputFixture = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func Test_parse(t *testing.T) {
	got := parse(inputFixture)

	want := heightmap{
		{'S', 'a', 'b', 'q', 'p', 'o', 'n', 'm'},
		{'a', 'b', 'c', 'r', 'y', 'x', 'x', 'l'},
		{'a', 'c', 'c', 's', 'z', 'E', 'x', 'k'},
		{'a', 'c', 'c', 't', 'u', 'v', 'w', 'j'},
		{'a', 'b', 'd', 'e', 'f', 'g', 'h', 'i'},
	}

	assert.Equal(t, want, got)
}

func Test_heightmap_height(t *testing.T) {
	f := parse(inputFixture)

	assert.Equal(t, 1, f.height(1, 0))
	assert.Equal(t, 2, f.height(2, 0))
	assert.Equal(t, 3, f.height(1, 2))
	assert.Equal(t, 26, f.height(4, 2))
}

func Test_heightmap_start(t *testing.T) {
	x, y := parse(inputFixture).start()

	assert.Equal(t, 0, x)
	assert.Equal(t, 0, y)
}

func Test_heightmap_end(t *testing.T) {
	x, y := parse(inputFixture).end()

	assert.Equal(t, 5, x)
	assert.Equal(t, 2, y)
}
