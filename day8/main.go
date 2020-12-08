package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var file = flag.String("file", "input.txt", "input file")

func main() {
	flag.Parse()
	var input []string

	file, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	pgm := parse(input)

	acc, _ := pgm.run()
	fmt.Println(acc)
	fmt.Println(pgm.runWithSelfHeal())
}

func parse(in []string) program {
	ins := regexp.MustCompile(`^(acc|nop|jmp) ([+-]\d+)$`)

	p := program{}
	for _, i := range in {
		m := ins.FindStringSubmatch(i)

		if len(m) != 3 {
			continue
		}

		c := m[1]
		v, _ := strconv.Atoi(m[2])
		p = append(p, instruction{c, v})
	}
	return p
}

type instruction struct {
	code  string
	value int
}

type program []instruction

func (p program) run() (int, error) {
	usd := make(map[int]bool)
	acc := 0
	pnt := 0

	for {
		if pnt == len(p) {
			return acc, nil
		}

		if pnt > len(p) {
			return acc, errors.New("pointer was out of bounds")
		}

		if _, exist := usd[pnt]; exist {
			return acc, errors.New("instruction ran twice")
		}

		ins := p[pnt]
		usd[pnt] = true

		switch ins.code {
		case "acc":
			acc += ins.value
			pnt++

		case "jmp":
			pnt += ins.value

		case "nop":
			pnt++
		}
	}
}

func (p program) runWithSelfHeal() int {
	cp := make(program, len(p))
	copy(cp, p)

	for i, ins := range p {
		if ins.code != "jmp" && ins.code != "nop" {
			continue
		}

		if ins.code == "jmp" {
			cp[i].code = "nop"
		} else {
			cp[i].code = "jmp"
		}

		acc, err := cp.run()

		if err != nil {
			copy(cp, p)
			continue
		}

		return acc
	}

	return -1
}
