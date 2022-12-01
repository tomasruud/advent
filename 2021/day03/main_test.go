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
		want tobogganMap
	}{
		{"parses", args{example}, tobogganMap{
			"..##.......",
			"#...#...#..",
			".#....#..#.",
			"..#.#...#.#",
			".#...##..#.",
			"..#.##.....",
			".#.#.#....#",
			".#........#",
			"#.##...#...",
			"#...##....#",
			".#..#...#.#",
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

func Test_tobogganMap_hasTreeAt(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		t       tobogganMap
		args    args
		want    bool
		wantErr bool
	}{
		{"finds tree", parse(example), args{3, 0}, true, false},
		{"notices when out of range", parse(example), args{3, 200}, false, true},
		{"finds not tree", parse(example), args{2, 2}, false, false},
		{"finds tree out of x bounds", parse(example), args{15, 5}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.hasTreeAt(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("hasTreeAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hasTreeAt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tobogganMap_countTreesInPath(t *testing.T) {
	type args struct {
		xStep int
		yStep int
	}
	tests := []struct {
		name string
		t    tobogganMap
		args args
		want int
	}{
		{"counts proper", parse(example), args{3, 1}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.countTreesInPath(tt.args.xStep, tt.args.yStep); got != tt.want {
				t.Errorf("countTreesInPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
