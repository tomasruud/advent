package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const part1fixture = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func Test_parse(t *testing.T) {
	got, err := parse(part1fixture)

	want := rucksacks{
		{
			compartmentA: []item("vJrwpWtwJgWr"),
			compartmentB: []item("hcsFMMfFFhFp"),
		},
		{
			compartmentA: []item("jqHRNqRjqzjGDLGL"),
			compartmentB: []item("rsFMfFZSrLrFZsSL"),
		},
		{
			compartmentA: []item("PmmdzqPrV"),
			compartmentB: []item("vPwwTWBwg"),
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, want, got[:3])
}

func Test_rucksack_union(t *testing.T) {
	fixture, err := parse(part1fixture)

	want := [][]item{
		{'p'},
		{'L'},
		{'P'},
		{'v'},
		{'t'},
		{'s'},
	}

	var got [][]item
	for _, i := range fixture {
		got = append(got, i.intersect())
	}

	assert.NoError(t, err)
	assert.Equal(t, want, got)

}

func Test_item_priority(t *testing.T) {
	tests := []struct {
		i    item
		want int
	}{
		{i: item('a'), want: 1},
		{i: item('b'), want: 2},
		{i: item('y'), want: 25},
		{i: item('z'), want: 26},
		{i: item('A'), want: 27},
		{i: item('B'), want: 28},
		{i: item('Y'), want: 51},
		{i: item('Z'), want: 52},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%c", tt.i), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.i.priority(), "priority()")
		})
	}
}

func Test_rucksacks_sumIntersections(t *testing.T) {
	fixture, err := parse(part1fixture)

	assert.NoError(t, err)
	assert.Equal(t, 157, fixture.sumIntersections())
}

func Test_group_badge(t *testing.T) {
	fixture := group{
		{
			compartmentA: []item("vJrwpWtwJgWr"),
			compartmentB: []item("hcsFMMfFFhFp"),
		},
		{
			compartmentA: []item("jqHRNqRjqzjGDLGL"),
			compartmentB: []item("rsFMfFZSrLrFZsSL"),
		},
		{
			compartmentA: []item("PmmdzqPrV"),
			compartmentB: []item("vPwwTWBwg"),
		},
	}

	assert.Equal(t, item('r'), fixture.badge())
}
