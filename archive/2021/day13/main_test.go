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
		name  string
		args  args
		want  timestamp
		want1 schedule
	}{
		{"parses", args{example}, 939, schedule{7, 13, 59, 31, 19}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parse(tt.args.in)
			if got != tt.want {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_schedule_nextBus(t *testing.T) {
	ts, sc := parse(example)

	type args struct {
		a timestamp
	}
	tests := []struct {
		name  string
		s     schedule
		args  args
		want  bus
		want1 timestamp
	}{
		{"gets next bus", sc, args{ts}, 59, 944},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.nextBus(tt.args.a)
			if got != tt.want {
				t.Errorf("nextBus() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("nextBus() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_schedule_next(t *testing.T) {
	ts, sc := parse(example)

	type args struct {
		a timestamp
	}
	tests := []struct {
		name string
		s    schedule
		args args
		want int
	}{
		{"calculates", sc, args{ts}, 295},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.next(tt.args.a); got != tt.want {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseV2(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want schedule
	}{
		{"parses", args{example}, schedule{7, 13, any, any, 59, any, 31, 19}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseV2(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_schedule_firstChain(t *testing.T) {
	type args struct {
		start timestamp
	}
	tests := []struct {
		name string
		s    schedule
		args args
		want timestamp
	}{
		{"finds chain", parseV2(example),args{1}, 1068781},
		{"finds chain", parseV2(example2),args{1}, 3417},
		{"finds chain", parseV2(example3),args{1}, 754018},
		{"finds chain", parseV2(example4),args{1}, 779210},
		{"finds chain", parseV2(example5), args{1},1261476},
		{"finds chain", parseV2(example6), args{1},1202161486},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.firstChain(tt.args.start); got != tt.want {
				t.Errorf("firstChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
