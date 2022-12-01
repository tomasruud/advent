package main

import (
	"reflect"
	"testing"
)

func Test_ship_manhattan(t *testing.T) {
	tests := []struct {
		name   string
		fields ship
		want   int
	}{
		{"calculates", ship{0, 0, east}.moves(parse(example1)), 25},
		{"calculates", ship{214, -72, east}, 286},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.manhattan(); got != tt.want {
				t.Errorf("manhattan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want []move
	}{
		{"parses", args{example1}, []move{
			{forward, 10},
			{north, 3},
			{forward, 7},
			{right, 90},
			{forward, 11},
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

func Test_ship_move(t *testing.T) {
	type args struct {
		m move
	}
	tests := []struct {
		name   string
		fields ship
		args   args
		want   ship
	}{
		{"move", ship{0, 0, east}, args{parse(example1)[0]}, ship{10, 0, east}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ship{
				x:         tt.fields.x,
				y:         tt.fields.y,
				direction: tt.fields.direction,
			}
			if got := s.move(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_waypointShip_move(t *testing.T) {
	type fields struct {
		x int
		y int
		s ship
	}
	type args struct {
		m move
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   waypoint
	}{
		{
			"calculates move",
			fields{10, 1, ship{0, 0, east}},
			args{move{forward, 10}},
			waypoint{10, 1, ship{100, 10, east}},
		},
		{
			"calculates move",
			fields{10, 1, ship{100, 10, east}},
			args{move{north, 3}},
			waypoint{10, 4, ship{100, 10, east}},
		},
		{
			"calculates move",
			fields{10, 4, ship{100, 10, east}},
			args{move{forward, 7}},
			waypoint{10, 4, ship{170, 38, east}},
		},
		{
			"calculates move",
			fields{10, 4, ship{170, 38, east}},
			args{move{right, 90}},
			waypoint{4, -10, ship{170, 38, east}},
		},
		{
			"calculates move",
			fields{4, -10, ship{170, 38, east}},
			args{move{forward, 11}},
			waypoint{4, -10, ship{214, -72, east}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := waypoint{
				x: tt.fields.x,
				y: tt.fields.y,
				s: tt.fields.s,
			}
			if got := w.move(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("move() = %v, want %v", got, tt.want)
			}
		})
	}
}
