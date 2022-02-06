package data_structure

import "fmt"

const MAX = 10

type Stack [MAX]int

var top int

func (s *Stack) Initialize() {
	top = 0
}

func (s *Stack) IsEmpty() bool {
	return top == 0
}

func (s *Stack) IsFull() bool {
	return top >= MAX - 1
}

func (s *Stack) Push(x int) error {
	if s.IsFull() {
		return fmt.Errorf("stack is full")
	}
	top++
	s[top] = x

	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	top--
	return s[top+1], nil
}