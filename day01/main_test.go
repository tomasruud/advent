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
		want report
	}{
		{"parses", args{example}, report{1721, 979, 366, 299, 675, 1456}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_report_double(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		r       report
		args    args
		want    int
		wantErr bool
	}{
		{"finds proper value", parse(example), args{2020}, 514579, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.double(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("double() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("double() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_report_triple(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		r       report
		args    args
		want    int
		wantErr bool
	}{
		{"finds proper value", parse(example), args{2020}, 241861950, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.triple(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("triple() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("triple() got = %v, want %v", got, tt.want)
			}
		})
	}
}