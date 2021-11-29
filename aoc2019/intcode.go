package aoc2019

import (
	"errors"
)

const (
	OP_ADD = 1
	OP_MUL = 2
	OP_END = 99
)

var (
	ErrInvalid = errors.New("Invalid Operation")
	EndProgram = errors.New("End of program")
)

type Intcode struct {
	ptr int
	mem []int
}

func NewIntcode (program []int) Intcode {
	progcopy := make([]int, len(program))
	copy(progcopy, program)
	return Intcode{0, progcopy}
}

func (c *Intcode) RunFull() error {
	for {
		switch err := c.step(); err {
		case nil:
			continue
		case EndProgram:
			return nil
		default:
			return err
		}
	}
}

func (c *Intcode) Get (a int) (int, error) {
	if a < 0 || a >= len(c.mem) {
		return 0, errors.New("Invalid address")
	}
	return c.mem[a], nil
}

func (c *Intcode) Set (a int, v int) error {
	if a < 0 || a >= len(c.mem) {
		return errors.New("Invalid address")
	}
	c.mem[a] = v
	return nil
}

func (c *Intcode) step () error {
	switch c.mem[c.ptr] {
	case OP_ADD:
		return c.add()
	case OP_MUL:
		return c.mul()
	case OP_END:
		return EndProgram
	default:
		return ErrInvalid
	}
}

func (c *Intcode) add () error {
	if c.mem[c.ptr] != OP_ADD {
		return ErrInvalid
	}
	a := c.mem[c.ptr + 1]
	b := c.mem[c.ptr + 2]
	d := c.mem[c.ptr + 3]
	c.mem[d] = c.mem[a] + c.mem[b]
	c.ptr += 4
	return nil
}

func (c *Intcode) mul () error {
	if c.mem[c.ptr] != OP_MUL {
		return ErrInvalid
	}
	a := c.mem[c.ptr + 1]
	b := c.mem[c.ptr + 2]
	d := c.mem[c.ptr + 3]
	c.mem[d] = c.mem[a] * c.mem[b]
	c.ptr += 4
	return nil
}

