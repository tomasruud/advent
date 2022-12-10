package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const inputFixture = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func Test_parse(t *testing.T) {
	got, got2, err := parse(inputFixture)

	want := ship{
		stacks: []crates{
			{"N", "Z"},
			{"D", "C", "M"},
			{"P"},
		},
	}

	want2 := moves{
		{amount: 1, from: 2, to: 1},
		{amount: 3, from: 1, to: 3},
		{amount: 2, from: 2, to: 1},
		{amount: 1, from: 1, to: 2},
	}

	assert.NoError(t, err)
	assert.Equal(t, want, got)
	assert.Equal(t, want2, got2)
}

func Test_ship_top(t *testing.T) {
	fixture := ship{
		stacks: []crates{
			{"C"},
			{"M"},
			{"Z", "N", "D", "P"},
		},
	}

	want := crates{"C", "M", "Z"}

	assert.Equal(t, want, fixture.top())
}

func Test_crateMover9000(t *testing.T) {
	s := &ship{
		stacks: []crates{
			{"N", "Z"},
			{"D", "C", "M"},
			{"P"},
		},
	}

	want := &ship{
		stacks: []crates{
			{"C"},
			{"M"},
			{"Z", "N", "D", "P"},
		},
	}

	err := crateMover9000(s, moves{
		{amount: 1, from: 2, to: 1},
		{amount: 3, from: 1, to: 3},
		{amount: 2, from: 2, to: 1},
		{amount: 1, from: 1, to: 2},
	})

	assert.NoError(t, err)
	assert.Equal(t, want, s)
}

func Test_crateMover9001(t *testing.T) {
	s := &ship{
		stacks: []crates{
			{"N", "Z"},
			{"D", "C", "M"},
			{"P"},
		},
	}

	want := &ship{
		stacks: []crates{
			{"M"},
			{"C"},
			{"D", "N", "Z", "P"},
		},
	}

	err := crateMover9001(s, moves{
		{amount: 1, from: 2, to: 1},
		{amount: 3, from: 1, to: 3},
		{amount: 2, from: 2, to: 1},
		{amount: 1, from: 1, to: 2},
	})

	assert.NoError(t, err)
	assert.Equal(t, want, s)
}
