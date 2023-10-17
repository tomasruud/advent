package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_elves_topCalories(t *testing.T) {
	t.Run("it returns top n calorie count", func(t *testing.T) {
		fixture := elves{
			elf{calories: []int{1}},
			elf{calories: []int{6}},
			elf{calories: []int{200}},
			elf{calories: []int{20}},
		}

		want := elves{
			elf{calories: []int{200}},
			elf{calories: []int{20}},
			elf{calories: []int{6}},
			elf{calories: []int{1}},
		}

		assert.Equal(t, want, fixture.topCalories(4))
	})

	t.Run("it returns partial top n calorie count", func(t *testing.T) {
		fixture := elves{
			elf{calories: []int{1}},
			elf{calories: []int{6}},
			elf{calories: []int{200}},
			elf{calories: []int{20}},
		}

		want := elves{
			elf{calories: []int{200}},
			elf{calories: []int{20}},
		}

		assert.Equal(t, want, fixture.topCalories(2))
	})
}
