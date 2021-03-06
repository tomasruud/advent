package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	pgm := parse(input)

	fmt.Println(pgm.run())
	fmt.Println(pgm.runWithSelfHeal())
}

func parse(in string) program {
	ins := regexp.MustCompile(`^(acc|nop|jmp) ([+-]\d+)$`)

	p := program{}
	for _, i := range strings.Split(in, "\n") {
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

func (p program) runWithSelfHeal() (int, error) {
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

		return acc, nil
	}

	return -1, errors.New("unable to self heal")
}
