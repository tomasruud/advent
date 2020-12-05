package main

import "testing"

func Test_boardingPass_row(t *testing.T) {
	type fields struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"finds row", fields{"FBFBBFFRLR"}, 44},
		{"finds row", fields{"BFFFBBFRRR"}, 70},
		{"finds row", fields{"FFFBBBFRRR"}, 14},
		{"finds row", fields{"BBFFBBFRLL"}, 102},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := boardingPass{
				seatCode: tt.fields.code,
				plane:    plane{8, 128},
			}
			if got := s.row(); got != tt.want {
				t.Errorf("row() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardingPass_col(t *testing.T) {
	type fields struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"finds col", fields{"FBFBBFFRLR"}, 5},
		{"finds col", fields{"BFFFBBFRRR"}, 7},
		{"finds col", fields{"FFFBBBFRRR"}, 7},
		{"finds col", fields{"BBFFBBFRLL"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := boardingPass{
				seatCode: tt.fields.code,
				plane:    plane{8, 128},
			}
			if got := s.col(); got != tt.want {
				t.Errorf("col() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardingPass_seatID(t *testing.T) {
	type fields struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"finds seatID", fields{"FBFBBFFRLR"}, 357},
		{"finds seatID", fields{"BFFFBBFRRR"}, 567},
		{"finds seatID", fields{"FFFBBBFRRR"}, 119},
		{"finds seatID", fields{"BBFFBBFRLL"}, 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := boardingPass{
				seatCode: tt.fields.code,
				plane:    plane{8, 128},
			}
			if got := s.seatID(); got != tt.want {
				t.Errorf("seatID() = %v, want %v", got, tt.want)
			}
		})
	}
}
