package main

import (
	"reflect"
	"testing"
)

var sample = []string{
	"abc",
	"",
	"a",
	"b",
	"c",
	"",
	"ab",
	"ac",
	"",
	"a",
	"a",
	"a",
	"a",
	"",
	"b",
}

func Test_parse(t *testing.T) {
	type args struct {
		i []string
	}
	tests := []struct {
		name string
		args args
		want flight
	}{
		{"parses properly", args{sample}, flight{
			group{form{"a", "b", "c"}},
			group{form{"a"}, form{"b"}, form{"c"}},
			group{form{"a", "b"}, form{"a", "c"}},
			group{form{"a"}, form{"a"}, form{"a"}, form{"a"}},
			group{form{"b"}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_group_unique(t *testing.T) {
	tests := []struct {
		name string
		g    group
		want int
	}{
		{"counts properly", parse(sample)[0], 3},
		{"counts properly", parse(sample)[1], 3},
		{"counts properly", parse(sample)[2], 3},
		{"counts properly", parse(sample)[3], 1},
		{"counts properly", parse(sample)[4], 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.unique(); got != tt.want {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flight_uniqueSum(t *testing.T) {
	tests := []struct {
		name string
		f    flight
		want int
	}{
		{"counts properly", parse(sample), 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.uniqueSum(); got != tt.want {
				t.Errorf("uniqueSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_group_all(t *testing.T) {
	tests := []struct {
		name string
		g    group
		want int
	}{
		{"counts properly", parse(sample)[0], 3},
		{"counts properly", parse(sample)[1], 0},
		{"counts properly", parse(sample)[2], 1},
		{"counts properly", parse(sample)[3], 1},
		{"counts properly", parse(sample)[4], 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.all(); got != tt.want {
				t.Errorf("all() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flight_allSum(t *testing.T) {
	tests := []struct {
		name string
		f    flight
		want int
	}{
		{"counts properly", parse(sample), 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.allSum(); got != tt.want {
				t.Errorf("allSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
