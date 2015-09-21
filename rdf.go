package main

import (
	"fmt"
)

const (
	TotalCpu = 9
	TotalMem = 18
)

type Task struct {
	user string
	cpu  int
	mem  int
	drs  float64
}

type Tasks []Task

func (ts Tasks) Len() int      {}
func (ts Tasks) Less() bool    {}
func (ts Tasks) Swap(i, j int) {}

func drf(ts []Task) {

}

func main() {
	a := Task{
		user: "user a",
		cpu:  1,
		mem:  4,
	}
	b := Task{
		user: "user b",
		cpu:  3,
		mem:  1,
	}
	tasks := []Task{a, b}
}
