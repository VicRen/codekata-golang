package main

type Stack []interface {
}

func MakeStack() *Stack {
	return &Stack{}
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s[len(s)-1]
}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	ret := (*s)[len(*s)-1]
	*s = (*s)[:(len(*s) - 1)]
	return ret
}
