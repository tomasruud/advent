package main

import (
	"reflect"
	"testing"
)

var sample = []string{
	"35",
	"20",
	"15",
	"25",
	"47",
	"40",
	"62",
	"55",
	"65",
	"95",
	"102",
	"117",
	"150",
	"182",
	"127",
	"219",
	"299",
	"277",
	"309",
	"576",
}

func Test_parse(t *testing.T) {
	type args struct {
		in       []string
		preamble int
	}
	tests := []struct {
		name string
		args args
		want xmas
	}{
		{"parses", args{sample, 5}, xmas{
			pre:  5,
			list: []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.in, tt.args.preamble); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xmas_invalid(t *testing.T) {
	tests := []struct {
		name   string
		fields xmas
		want   int
	}{
		{"finds invalid", parse(sample, 5), 127},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := xmas{
				pre:  tt.fields.pre,
				list: tt.fields.list,
			}
			if got := x.invalid(); got != tt.want {
				t.Errorf("invalid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xmas_valid(t *testing.T) {
	type args struct {
		v int
		i int
	}
	tests := []struct {
		name   string
		fields xmas
		args   args
		want   bool
	}{
		{"accepts valid", parse(sample, 5), args{102, 10}, true},
		{"accepts valid", parse(sample, 5), args{65, 8}, true},
		{"accepts valid", parse(sample, 5), args{95, 9}, true},
		{"fails invalid", parse(sample, 5), args{127, 14}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := xmas{
				pre:  tt.fields.pre,
				list: tt.fields.list,
			}
			if got := x.valid(tt.args.v, tt.args.i); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xmas_weak(t *testing.T) {
	tests := []struct {
		name   string
		fields xmas
		want   int
	}{
		{"finds weak", parse(sample, 5), 62},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := xmas{
				pre:  tt.fields.pre,
				list: tt.fields.list,
			}
			if got := x.weak(); got != tt.want {
				t.Errorf("weak() = %v, want %v", got, tt.want)
			}
		})
	}
}
