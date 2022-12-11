package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const inputFixture = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

func Test(t *testing.T) {
	t.Run("it parses input", func(t *testing.T) {
		got, err := parse(inputFixture)

		want := monkeys{
			{worries: []int{79, 98}, b: p(19), op: "*", mod: 23, yes: 2, no: 3},
			{worries: []int{54, 65, 75, 74}, b: p(6), op: "+", mod: 19, yes: 2, no: 0},
			{worries: []int{79, 60, 97}, op: "*", mod: 13, yes: 1, no: 3},
			{worries: []int{74}, b: p(3), op: "+", mod: 17, yes: 0, no: 1},
		}

		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("it works with relief factor 3", func(t *testing.T) {
		f, _ := parse(inputFixture)

		f.rounds(20, 3)

		assert.Equal(t, 10605, f.monkeyBusiness())
	})

	t.Run("it works with relief factor 1", func(t *testing.T) {
		f, _ := parse(inputFixture)

		f.rounds(10000, 1)

		assert.Equal(t, 2713310158, f.monkeyBusiness())
	})

	t.Run("it works for 1000 rounds with relief 1", func(t *testing.T) {
		f, _ := parse(inputFixture)

		f.rounds(1000, 1)

		want := []int{5204, 4792, 199, 5192}

		for i, w := range want {
			assert.Equal(t, w, f[i].inspected)
		}
	})

	t.Run("it works for 3000 rounds with relief 1", func(t *testing.T) {
		f, _ := parse(inputFixture)

		f.rounds(3000, 1)

		want := []int{15638, 14358, 587, 15593}

		for i, w := range want {
			assert.Equal(t, w, f[i].inspected)
		}
	})

	t.Run("it works for 6000 rounds with relief 1", func(t *testing.T) {
		f, _ := parse(inputFixture)

		f.rounds(6000, 1)

		want := []int{31294, 28702, 1165, 31204}

		for i, w := range want {
			assert.Equal(t, w, f[i].inspected)
		}
	})
}

func p[T any](v T) *T {
	return &v
}
