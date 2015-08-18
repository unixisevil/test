package main

import (
	"errors"
)

type Stack []string

func (stack *Stack) Pop() (string, error) {
	tmp := *stack
	if len(tmp) == 0 {
		return "", errors.New("can't Pop() an empty stack")
	}
	x := tmp[len(tmp)-1]
	*stack = tmp[:len(tmp)-1]
	return x, nil
}

func (stack *Stack) Push(x string) {
	*stack = append(*stack, x)
}

func (stack Stack) Top() (string, error) {
	if len(stack) == 0 {
		return "", errors.New("can't Top() an empty stack")
	}
	return stack[len(stack)-1], nil
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
