package main

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want bag
	}{
		{"parses", args{example1}, bag{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bag_sort(t *testing.T) {
	tests := []struct {
		name string
		b    bag
		want bag
	}{
		{"sorts", parse(example1), bag{1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bag_diffSize(t *testing.T) {
	tests := []struct {
		name string
		b    bag
		want int
	}{
		{"finds correct", parse(example1), 7*5},
		{"finds correct", parse(example2), 22*10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.diffSize(); got != tt.want {
				t.Errorf("diffSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_adapter_diff(t *testing.T) {
	type args struct {
		o adapter
	}
	tests := []struct {
		name string
		a    adapter
		args args
		want int
	}{
		{"diffs", adapter(5), args{9}, 4},
		{"diffs", adapter(9), args{5}, -4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.diff(tt.args.o); got != tt.want {
				t.Errorf("diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_adapter_connects(t *testing.T) {
	type args struct {
		o adapter
	}
	tests := []struct {
		name string
		a    adapter
		args args
		want bool
	}{
		{"passes valid", adapter(4), args{6}, true},
		{"passes valid", adapter(4), args{7}, true},
		{"fails invalid", adapter(4), args{8}, false},
		{"fails invalid", adapter(4), args{3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.connects(tt.args.o); got != tt.want {
				t.Errorf("connects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bag_chain(t *testing.T) {
	tests := []struct {
		name string
		b    bag
		want bag
	}{
		{"chains", parse(example1), bag{0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.chain(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bag_branches(t *testing.T) {
	tests := []struct {
		name string
		b    bag
		want int
	}{
		{"calculated properly", parse(example1), 8},
		{"calculated properly", parse(example2), 19208},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.branches(); got != tt.want {
				t.Errorf("branches() = %v, want %v", got, tt.want)
			}
		})
	}
}