package main

import "errors"

type Stack []interface {
}

func NewStack() Stack {
	return Stack{}
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Push(item interface{}) error {
	*s = append(*s, item)
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret, nil
}

func (s Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s[len(s)-1]
}

func (s Stack) Len() int {
	return len(s)
}
