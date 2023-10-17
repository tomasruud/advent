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
		want list
	}{
		{"parses", args{example}, list{
			{1, 3, "a", "abcde"},
			{1, 3, "b", "cdefg"},
			{2, 9, "c", "ccccccccc"},
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

func Test_pw_valid(t *testing.T) {
	type fields struct {
		min  int
		max  int
		char string
		pass string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"passes valid", fields{1,3, "a", "abcde"}, true},
		{"faild invalid", fields{1,3, "b", "cdefg"}, false},
		{"passes valid", fields{2,9, "c", "ccccccccc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pw{
				min:  tt.fields.min,
				max:  tt.fields.max,
				char: tt.fields.char,
				pass: tt.fields.pass,
			}
			if got := p.valid(); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pw_validAtToboggan(t *testing.T) {
	type fields struct {
		min  int
		max  int
		char string
		pass string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"passes valid", fields{1,3, "a", "abcde"}, true},
		{"fail invalid", fields{1,3, "b", "cdefg"}, false},
		{"fail invalid", fields{2,9, "c", "ccccccccc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pw{
				min:  tt.fields.min,
				max:  tt.fields.max,
				char: tt.fields.char,
				pass: tt.fields.pass,
			}
			if got := p.validAtToboggan(); got != tt.want {
				t.Errorf("validAtToboggan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_valid(t *testing.T) {
	tests := []struct {
		name string
		ps   list
		want int
	}{
		{"counts valid", parse(example), 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.valid(); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_validAtToboggan(t *testing.T) {
	tests := []struct {
		name string
		ps   list
		want int
	}{
		{"counts valid", parse(example), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.validAtToboggan(); got != tt.want {
				t.Errorf("validAtToboggan() = %v, want %v", got, tt.want)
			}
		})
	}
}