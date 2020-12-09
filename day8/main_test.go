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
		want program
	}{
		{"parses", args{example}, program{
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
		{"runs properly", parse(example), 5, true},
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

func Test_program_runWithSelfHeal1(t *testing.T) {
	tests := []struct {
		name    string
		p       program
		want    int
		wantErr bool
	}{
		{"heals properly", parse(example), 8, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.runWithSelfHeal()
			if (err != nil) != tt.wantErr {
				t.Errorf("runWithSelfHeal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("runWithSelfHeal() got = %v, want %v", got, tt.want)
			}
		})
	}
}
