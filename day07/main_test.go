package main

import (
	"reflect"
	"testing"
)

func Test_bags_contain(t *testing.T) {
	type args struct {
		c kind
	}
	tests := []struct {
		name string
		bs   bags
		args args
		want int
	}{
		{"finds proper amount", parse(example1), args{"shiny gold"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bs.contains(tt.args.c); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bag_size(t *testing.T) {
	type fields struct {
		kind kind
		bags bags
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"finds size", fields{"1", newBags(newBag("2"), newBag("2"), newBag("3"))}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bag{
				kind: tt.fields.kind,
				bags: tt.fields.bags,
			}
			if got := b.size(); got != tt.want {
				t.Errorf("size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bags_find(t *testing.T) {
	type args struct {
		k kind
	}
	tests := []struct {
		name string
		bs   bags
		args args
		want *bag
	}{
		{"finds", parse(example1), args{"shiny gold"}, newBag("shiny gold")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bs.find(tt.args.k); !reflect.DeepEqual(&got.kind, &tt.want.kind) {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}
