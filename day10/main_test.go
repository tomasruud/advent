package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parse(t *testing.T) {
	f := `noop
addx 3
addx -5`

	want := []instruction{
		noop(0),
		addx(3),
		addx(-5),
	}

	got, err := parse(f)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
