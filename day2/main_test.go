package main

import (
	"reflect"
	"testing"
)

func Test_parseEntry(t *testing.T) {
	type args struct {
		e string
	}
	tests := []struct {
		name    string
		args    args
		want    pw
		wantErr bool
	}{
		{"creates from valid", args{e: "1-3 a: abcde"}, pw{1, 3, "a", "abcde"}, false},
		{"creates from valid", args{e: "1-3 b: cdefg"}, pw{1, 3, "b", "cdefg"}, false},
		{"creates from valid", args{e: "12-19 c: ccccccccc"}, pw{12, 19, "c", "ccccccccc"}, false},
		{"invalid gives err", args{e: "gdfgewqr323"}, pw{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseEntry(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseEntry() got = %v, want %v", got, tt.want)
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