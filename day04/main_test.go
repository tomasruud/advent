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
		want collection
	}{
		{"parses values properly", args{example}, collection{
			passport{ecl: "gry", pid: "860033327", eyr: "2020", hcl: "#fffffd", byr: "1937", iyr: "2017", cid: "147", hgt: "183cm"},
			passport{iyr: "2013", ecl: "amb", cid: "350", eyr: "2023", pid: "028048884", hcl: "#cfa07d", byr: "1929"},
			passport{hcl: "#ae17e1", iyr: "2013", eyr: "2024", ecl: "brn", pid: "760753108", byr: "1931", hgt: "179cm"},
			passport{hcl: "#cfa07d", eyr: "2025", pid: "166559648", iyr: "2011", ecl: "brn", hgt: "59in"},
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

func Test_collection_validCount(t *testing.T) {
	tests := []struct {
		name string
		pc   collection
		want int
	}{
		{"counts properly", parse(example), 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.validCount(); got != tt.want {
				t.Errorf("validCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passport_valid(t *testing.T) {
	tests := []struct {
		name   string
		fields passport
		want   bool
	}{
		{"passes valid", parse(example)[0], true},
		{"fails invalid", parse(example)[1], false},
		{"passes valid", parse(example)[2], true},
		{"fails invalid", parse(example)[3], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := passport{
				byr: tt.fields.byr,
				iyr: tt.fields.iyr,
				eyr: tt.fields.eyr,
				hgt: tt.fields.hgt,
				hcl: tt.fields.hcl,
				ecl: tt.fields.ecl,
				pid: tt.fields.pid,
				cid: tt.fields.cid,
			}
			if got := p.valid(); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passport_strictValid(t *testing.T) {
	tests := []struct {
		name   string
		fields passport
		want   bool
	}{
		{"passes valid 1", parse(valid)[0], true},
		{"passes valid 2", parse(valid)[1], true},
		{"passes valid 3", parse(valid)[2], true},
		{"passes valid 4", parse(valid)[3], true},
		{"fails invalid 1", parse(invalid)[0], false},
		{"fails invalid 2", parse(invalid)[1], false},
		{"fails invalid 3", parse(invalid)[2], false},
		{"fails invalid 4", parse(invalid)[3], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := passport{
				byr: tt.fields.byr,
				iyr: tt.fields.iyr,
				eyr: tt.fields.eyr,
				hgt: tt.fields.hgt,
				hcl: tt.fields.hcl,
				ecl: tt.fields.ecl,
				pid: tt.fields.pid,
				cid: tt.fields.cid,
			}
			if got := p.strictValid(); got != tt.want {
				t.Errorf("strictValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
