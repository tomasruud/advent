package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stream(t *testing.T) {
	tests := []struct {
		data   data
		marker marker
		want   int
	}{
		{data: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", marker: packet, want: 7},
		{data: "bvwbjplbgvbhsrlpgdmjqwftvncz", marker: packet, want: 5},
		{data: "nppdvjthqldpwncqszvftbrmjlhg", marker: packet, want: 6},
		{data: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", marker: packet, want: 10},
		{data: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", marker: packet, want: 11},
		{data: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", marker: message, want: 19},
		{data: "bvwbjplbgvbhsrlpgdmjqwftvncz", marker: message, want: 23},
		{data: "nppdvjthqldpwncqszvftbrmjlhg", marker: message, want: 23},
		{data: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", marker: message, want: 29},
		{data: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", marker: message, want: 26},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
			d := stream{data: tt.data}

			got, err := d.indexOf(tt.marker)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
