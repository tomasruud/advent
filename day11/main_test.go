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
		want grid
	}{
		{"parses", args{example}, grid{
			{"L", ".", "L", "L", ".", "L", "L", ".", "L", "L"},
			{"L", "L", "L", "L", "L", "L", "L", ".", "L", "L"},
			{"L", ".", "L", ".", "L", ".", ".", "L", ".", "."},
			{"L", "L", "L", "L", ".", "L", "L", ".", "L", "L"},
			{"L", ".", "L", "L", ".", "L", "L", ".", "L", "L"},
			{"L", ".", "L", "L", "L", "L", "L", ".", "L", "L"},
			{".", ".", "L", ".", "L", ".", ".", ".", ".", "."},
			{"L", "L", "L", "L", "L", "L", "L", "L", "L", "L"},
			{"L", ".", "L", "L", "L", "L", "L", "L", ".", "L"},
			{"L", ".", "L", "L", "L", "L", "L", ".", "L", "L"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_occupied(t *testing.T) {
	tests := []struct {
		name string
		g    grid
		want int
	}{
		{"counts occupied", parse(result5), 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.occupied(); got != tt.want {
				t.Errorf("occupied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_copy(t *testing.T) {
	tests := []struct {
		name string
		g    grid
		want grid
	}{
		{"makes copy", parse(example), parse(example)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_simulate(t *testing.T) {
	tests := []struct {
		name string
		g    grid
		want grid
	}{
		{"simulates", parse(example), parse(result1)},
		{"simulates", parse(result1), parse(result2)},
		{"simulates", parse(result2), parse(result3)},
		{"simulates", parse(result3), parse(result4)},
		{"simulates", parse(result4), parse(result5)},
		{"simulates", parse(result5), parse(result5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.simulate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simulate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_adjacent(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		g    grid
		args args
		want grid
	}{
		{"finds adjacent", parse(example), args{0, 0}, grid{{floor, empty, empty}}},
		{"finds adjacent", parse(example), args{1, 1}, grid{{empty, floor, empty, empty, empty, empty, floor, empty}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.adjacent(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("adjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_simulateOccupied(t *testing.T) {
	tests := []struct {
		name string
		g    grid
		want int
	}{
		{"simulates properly", parse(example), 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.simulateOccupied(); got != tt.want {
				t.Errorf("simulateOccupied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_closest(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		g    grid
		args args
		want grid
	}{
		{"finds closest", parse(example2), args{3, 4}, grid{{occupied, occupied, occupied, occupied, occupied, occupied, occupied, occupied}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.closest(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_simulateV2(t *testing.T) {
	tests := []struct {
		name string
		g    grid
		want grid
	}{
		{"simulates", parse(example3), parse(example3result)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.simulateV2(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simulateV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_simulateOccupiedV2(t *testing.T) {
	tests := []struct {
		name string
		g    grid
		want int
	}{
		{"works", parse(example3), 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.simulateOccupiedV2(); got != tt.want {
				t.Errorf("simulateOccupiedV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
