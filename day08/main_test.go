package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const inputFixture = `30373
25512
65332
33549
35390`

var fixture = forest{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func Test_parse(t *testing.T) {
	assert.Equal(t, fixture, parse(inputFixture))
}

func Test_forest_sumVisibleFromEdge(t *testing.T) {
	assert.Equal(t, 21, fixture.sumVisibleFromEdge())
}

func Test_forest_visibleFromEdge(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want bool
	}{
		{0, 0, true},
		{0, 1, true},
		{1, 0, true},
		{2, 3, true},
		{3, 2, true},
		{3, 2, true},
		{2, 2, false},
		{1, 3, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("[%d,%d]", tt.x, tt.y), func(t *testing.T) {
			assert.Equal(t, tt.want, fixture.visibleFromEdge(tt.x, tt.y))
		})
	}
}

func Test_forest_scenicScore(t *testing.T) {
	assert.Equal(t, 4, fixture.scenicScore(2, 1))
	assert.Equal(t, 8, fixture.scenicScore(2, 3))
}

func Test_forest_viewDistance(t *testing.T) {
	tests := []struct {
		x     int
		y     int
		d     direction
		want  int
		want2 bool
	}{
		{2, 1, up, 1, false},
		{2, 1, left, 1, true},
		{2, 1, right, 2, false},
		{2, 1, down, 2, true},
		{2, 3, up, 2, true},
		{2, 3, left, 2, false},
		{2, 3, down, 1, false},
		{2, 3, right, 2, true},
		{0, 0, left, 0, false},
		{0, 0, up, 0, false},
		{0, 0, down, 2, true},
		{0, 0, right, 2, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("[%d,%d] %s", tt.x, tt.y, tt.d), func(t *testing.T) {
			got, got2 := fixture.viewDistance(tt.d, tt.x, tt.y)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want2, got2)
		})
	}
}

func Test_forest_maxScenicScore(t *testing.T) {
	assert.Equal(t, 8, fixture.maxScenicScore())
}
