package main

import (
	"fmt"
)

func main() {
	s := stream{
		data: input,
	}

	p, err := s.indexOf(packet)
	if err != nil {
		panic(err)
	}

	fmt.Println("first packet found after", p)

	m, err := s.indexOf(message)
	if err != nil {
		panic(err)
	}

	fmt.Println("first message found after", m)
}

type stream struct {
	data data
}

func (s stream) indexOf(m marker) (int, error) {
	if len(s.data) < m.size {
		return -1, fmt.Errorf("data is shorter than marker size")
	}

	for i := 0; i < len(s.data)-m.size; i++ {
		if s.data[i:i+m.size].uniqueChars() != m.size {
			continue
		}

		return i + m.size, nil
	}

	return -1, fmt.Errorf("no marker start was found")
}

type data string

func (d data) uniqueChars() int {
	u := map[rune]bool{}
	for _, c := range d {
		if _, ok := u[c]; ok {
			continue
		}

		u[c] = true
	}
	return len(u)
}

type marker struct {
	size int
}

var (
	packet  = marker{size: 4}
	message = marker{size: 14}
)
