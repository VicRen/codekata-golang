package main

import (
	"fmt"
	"testing"
)

func reorderList(h *node) {
	printList(h)
	f, s := h, h
	for f.next != nil && f.next.next != nil {
		f = f.next.next
		s = s.next
	}
	sh := s.next
	s.next = nil
	printList(h)
	printList(sh)

	var pre *node
	for sh != nil {
		next := sh.next
		sh.next = pre
		pre = sh
		sh = next
	}
	printList(pre)

	p1, p2 := h, pre
	for p2.next != nil {
		next1, next2 := p1.next, p2.next
		p1.next = p2
		p2.next = next1
		p1, p2 = next1, next2
	}
	printList(h)
}

func TestReorderList(t *testing.T) {
	reorderList(makeList(19))
}

func printList(h *node) {
	for h != nil {
		fmt.Printf("%d->", h.value)
		h = h.next
	}
	fmt.Println("nil")
}

func makeList(count int) *node {
	var h *node
	var p *node
	for i := 0; i < count; i++ {
		n := &node{
			value: i,
		}
		if h == nil {
			h = n
			p = h
		}
		p.next = n
		p = n
	}
	return h
}

type node struct {
	next  *node
	value int
}
