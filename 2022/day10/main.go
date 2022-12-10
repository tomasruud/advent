package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	is, err := parse(input)
	if err != nil {
		panic(err)
	}

	trace := trace{}

	crt := crt{
		width:  40,
		height: 6,
		cursor: 1,
	}

	cpu := &cpu{
		x: 1,
		listeners: map[event][]listener{
			clockCycle: {
				trace.onClockCycle,
				crt.onClockCycle,
			},
			instructionDone: {
				crt.onInstructionDone,
			},
		},
	}

	cpu.run(is...)

	fmt.Println("trace sum", trace.sum)
	fmt.Println("\noutput drawing:")
	crt.flush()
}

func parse(in string) ([]instruction, error) {
	var is []instruction
	for _, ln := range strings.Split(in, "\n") {
		ln = strings.TrimSpace(ln)

		if ln == "noop" {
			is = append(is, noop(0))
		} else if strings.HasPrefix(ln, "addx") {
			ln = strings.TrimPrefix(ln, "addx ")

			n, err := strconv.Atoi(ln)
			if err != nil {
				return is, fmt.Errorf("unable to parse line: %w", err)
			}

			is = append(is, addx(n))
		}
	}
	return is, nil
}

type trace struct {
	sum int
}

func (t *trace) onClockCycle(cpu *cpu) {
	if cpu.cycle < 20 {
		return
	}

	if (cpu.cycle-20)%40 == 0 {
		t.sum += cpu.signalStrength()
	}
}

type crt struct {
	width  int
	height int
	cursor int
	buffer string
}

func (c *crt) flush() {
	for row := 0; row < c.height; row++ {
		start := row * c.width
		if start >= len(c.buffer) {
			break
		}

		end := (row + 1) * c.width
		if end >= len(c.buffer) {
			end = len(c.buffer)
		}

		fmt.Println(c.buffer[start:end])
	}
	c.buffer = ""
}

func (c *crt) draw(char rune) {
	c.buffer = c.buffer + string(char)
}

func (c *crt) onClockCycle(cpu *cpu) {
	x := (cpu.cycle - 1) % c.width
	if x == c.cursor-1 || x == c.cursor || x == c.cursor+1 {
		c.draw('#')
	} else {
		c.draw('.')
	}
}

func (c *crt) onInstructionDone(cpu *cpu) {
	c.cursor = cpu.x
}

type cpu struct {
	cycle     int
	x         int
	listeners map[event][]listener
}

func (c *cpu) signalStrength() int {
	return c.cycle * c.x
}

func (c *cpu) run(ins ...instruction) {
	for _, in := range ins {
		for i := 0; i < in.cycles(); i++ {
			c.cycle++
			c.emit(clockCycle)
		}
		in.run(c)
		c.emit(instructionDone)
	}
}

func (c *cpu) emit(e event) {
	if c.listeners == nil {
		return
	}

	for _, l := range c.listeners[e] {
		l(c)
	}
}

type listener func(cpu *cpu)

type event int

const (
	clockCycle event = iota
	instructionDone
)

type instruction interface {
	cycles() int
	run(*cpu)
}

type noop int

func (n noop) cycles() int {
	return 1
}

func (n noop) run(_ *cpu) {
}

type addx int

func (a addx) cycles() int {
	return 2
}

func (a addx) run(c *cpu) {
	c.x += int(a)
}
