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
		want flight
	}{
		{"parses", args{example}, flight{
			plane{8, 128},
			[]boardingPass{
				{"FBFBBFFRLR"},
				{"BFFFBBFRRR"},
				{"FFFBBBFRRR"},
				{"BBFFBBFRLL"},
			},
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

func Test_plane_row(t *testing.T) {
	type fields struct {
		cols int
		rows int
	}
	type args struct {
		b boardingPass
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"finds row", fields{8, 128}, args{parse(example).passes[0]}, 44},
		{"finds row", fields{8, 128}, args{parse(example).passes[1]}, 70},
		{"finds row", fields{8, 128}, args{parse(example).passes[2]}, 14},
		{"finds row", fields{8, 128}, args{parse(example).passes[3]}, 102},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := plane{
				cols: tt.fields.cols,
				rows: tt.fields.rows,
			}
			if got := p.row(tt.args.b); got != tt.want {
				t.Errorf("row() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_plane_col(t *testing.T) {
	type fields struct {
		cols int
		rows int
	}
	type args struct {
		b boardingPass
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"finds col", fields{8, 128}, args{parse(example).passes[0]}, 5},
		{"finds col", fields{8, 128}, args{parse(example).passes[1]}, 7},
		{"finds col", fields{8, 128}, args{parse(example).passes[2]}, 7},
		{"finds col", fields{8, 128}, args{parse(example).passes[3]}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := plane{
				cols: tt.fields.cols,
				rows: tt.fields.rows,
			}
			if got := p.col(tt.args.b); got != tt.want {
				t.Errorf("col() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_plane_seatID(t *testing.T) {
	type fields struct {
		cols int
		rows int
	}
	type args struct {
		b boardingPass
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"finds id", fields{8, 128}, args{parse(example).passes[0]}, 357},
		{"finds id", fields{8, 128}, args{parse(example).passes[1]}, 567},
		{"finds id", fields{8, 128}, args{parse(example).passes[2]}, 119},
		{"finds id", fields{8, 128}, args{parse(example).passes[3]}, 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := plane{
				cols: tt.fields.cols,
				rows: tt.fields.rows,
			}
			if got := p.seatID(tt.args.b); got != tt.want {
				t.Errorf("seatID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flight_maxSeatID(t *testing.T) {
	type fields struct {
		plane  plane
		passes []boardingPass
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"finds max", fields(parse(example)), 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := flight{
				plane:  tt.fields.plane,
				passes: tt.fields.passes,
			}
			if got := f.maxSeatID(); got != tt.want {
				t.Errorf("maxSeatID() = %v, want %v", got, tt.want)
			}
		})
	}
}