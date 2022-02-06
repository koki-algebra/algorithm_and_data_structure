package main

import (
	"algorithm_and_data_structure/data_structure"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	// sc.Split(bufio.ScanWords)

	var (
		s data_structure.Stack
		// inputs []string
	)

	// Stack initialize
	s.Initialize()

	if sc.Scan() {
		for _, c := range strings.Split(sc.Text(), " ") {
			switch c {
			case "+":
				x, err := s.Pop()
				if err != nil {
					panic(err)
				}
				y, err := s.Pop()
				if err != nil {
					panic(err)
				}
				s.Push(y+x)
			case "-":
				x, err := s.Pop()
				if err != nil {
					panic(err)
				}
				y, err := s.Pop()
				if err != nil {
					panic(err)
				}
				s.Push(y-x)
			case "*":
				x, err := s.Pop()
				if err != nil {
					panic(err)
				}
				y, err := s.Pop()
				if err != nil {
					panic(err)
				}
				s.Push(y*x)
			default:
				n, err := strconv.Atoi(c)
				if err != nil {
					panic(err)
				}
				if err := s.Push(n); err != nil {
					panic(err)
				}
			}
		}
	}

	x, err := s.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Println(x)
}