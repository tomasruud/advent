package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const fixture = `A Y
B X
C Z`

func Test_parsePart1(t *testing.T) {
	want := guide{
		{opponent: rock, me: paper},
		{opponent: paper, me: rock},
		{opponent: scissor, me: scissor},
	}

	got, err := parsePart1(fixture)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func Test_parsePart2(t *testing.T) {
	want := guide{
		{opponent: rock, me: rock},
		{opponent: paper, me: rock},
		{opponent: scissor, me: rock},
	}

	got, err := parsePart2(fixture)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func Test_guide_score(t *testing.T) {
	g := guide{
		{opponent: rock, me: paper},
		{opponent: paper, me: paper},
		{opponent: scissor, me: paper},
	}

	assert.Equal(t, 15, g.score())
}
