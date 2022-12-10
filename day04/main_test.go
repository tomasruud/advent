package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const fixturePart1 = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func Test_parse(t *testing.T) {
	want := pairs{
		{
			{sections: []section{2, 3, 4}},
			{sections: []section{6, 7, 8}},
		},
		{
			{sections: []section{2, 3}},
			{sections: []section{4, 5}},
		},
		{
			{sections: []section{5, 6, 7}},
			{sections: []section{7, 8, 9}},
		},
		{
			{sections: []section{2, 3, 4, 5, 6, 7, 8}},
			{sections: []section{3, 4, 5, 6, 7}},
		},
		{
			{sections: []section{6}},
			{sections: []section{4, 5, 6}},
		},
		{
			{sections: []section{2, 3, 4, 5, 6}},
			{sections: []section{4, 5, 6, 7, 8}},
		},
	}

	got, err := parse(fixturePart1)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func Test_elf_hasSection(t *testing.T) {
	fixture := elf{
		sections: []section{1, 2, 3, 4, 5},
	}

	got := fixture.hasSection([]section{2, 3, 4})

	assert.True(t, got)
}

func Test_pair_fullyContains(t *testing.T) {
	t.Run("it handles a->b", func(t *testing.T) {
		fixture := pair{
			{sections: []section{2, 3, 4, 5, 6, 7, 8}},
			{sections: []section{3, 4, 5, 6, 7}},
		}

		assert.True(t, fixture.fullyContains())
	})

	t.Run("it handles b->a", func(t *testing.T) {
		fixture := pair{
			{sections: []section{3, 4, 5, 6, 7}},
			{sections: []section{2, 3, 4, 5, 6, 7, 8}},
		}

		assert.True(t, fixture.fullyContains())
	})
}

func Test_elf_overlaps(t *testing.T) {
	a := elf{sections: []section{1, 2, 3}}
	b := elf{sections: []section{3, 4, 5}}

	assert.True(t, a.overlaps(b))
}

func Test_pair_overlaps(t *testing.T) {
	fixture := pair{
		{sections: []section{1, 2, 3}},
		{sections: []section{3, 4, 5}},
	}

	assert.True(t, fixture.overlaps())
}
