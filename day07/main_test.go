package main

import (
	"reflect"
	"testing"
)

var sample = []string{
	"light red bags contain 1 bright white bag, 2 muted yellow bags.",
	"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
	"bright white bags contain 1 shiny gold bag.",
	"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
	"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
	"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
	"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
	"faded blue bags contain no other bags.",
	"dotted black bags contain no other bags.",
}

var sample2 = []string{
	"shiny gold bags contain 2 dark red bags.",
	"dark red bags contain 2 dark orange bags.",
	"dark orange bags contain 2 dark yellow bags.",
	"dark yellow bags contain 2 dark green bags.",
	"dark green bags contain 2 dark blue bags.",
	"dark blue bags contain 2 dark violet bags.",
	"dark violet bags contain no other bags.",
}

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
		{"finds proper amount", parse(sample), args{"shiny gold"}, 4},
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
		{"finds", parse(sample), args{"shiny gold"}, newBag("shiny gold")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bs.find(tt.args.k); !reflect.DeepEqual(&got.kind, &tt.want.kind) {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}
