package main

import (
	"reflect"
	"testing"
)

var sample = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"jmp -4",
	"acc +6",
}

func Test_parse(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want program
	}{
		{"parses", args{sample}, program{
			{"nop", 0},
			{"acc", 1},
			{"jmp", 4},
			{"acc", 3},
			{"jmp", -3},
			{"acc", -99},
			{"acc", 1},
			{"jmp", -4},
			{"acc", 6},
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

func Test_program_run(t *testing.T) {
	tests := []struct {
		name    string
		p       program
		want    int
		wantErr bool
	}{
		{"runs properly", parse(sample), 5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.run()
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("run() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_program_runWithSelfHeal(t *testing.T) {
	tests := []struct {
		name string
		p    program
		want int
	}{
		{"heals properly", parse(sample), 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.runWithSelfHeal(); got != tt.want {
				t.Errorf("runWithSelfHeal() = %v, want %v", got, tt.want)
			}
		})
	}
}
