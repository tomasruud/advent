package main

import "testing"

var input = []string{
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
}

func Test_tobogganMap_hasTreeAt(t1 *testing.T) {
	type fields struct {
		lines []string
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{"finds tree", fields{input}, args{3, 0}, true, false},
		{"notices when out of range", fields{input}, args{3, 200}, false, true},
		{"finds not tree", fields{input}, args{2, 2}, false, false},
		{"finds tree out of x bounds", fields{input}, args{15, 5}, true, false},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := tobogganMap{
				lines: tt.fields.lines,
			}
			got, err := t.hasTreeAt(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t1.Errorf("hasTreeAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("hasTreeAt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tobogganMap_countTreesInPath(t1 *testing.T) {
	type fields struct {
		lines []string
	}
	type args struct {
		xStep int
		yStep int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"counts proper", fields{input}, args{3, 1}, 7},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := tobogganMap{
				lines: tt.fields.lines,
			}
			if got := t.countTreesInPath(tt.args.xStep, tt.args.yStep); got != tt.want {
				t1.Errorf("countTreesInPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
