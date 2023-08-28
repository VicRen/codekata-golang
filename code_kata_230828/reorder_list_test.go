package main

import (
	"fmt"
	"testing"
)

func ReorderList(head *node) {
	printList(head)
	f, s := head, head
	for f.next != nil && f.next.next != nil {
		f = f.next.next
		s = s.next
	}
	sh := s.next
	s.next = nil
	printList(head)
	printList(sh)

	var pre *node
	for sh != nil {
		next := sh.next
		sh.next = pre
		pre = sh
		sh = next
	}
	printList(pre)

	p1, p2 := head, pre
	for p2 != nil {
		next1, next2 := p1.next, p2.next
		p1.next = p2
		p2.next = next1
		p1, p2 = next1, next2
	}

	printList(head)
}

func printList(n *node) {
	for n != nil {
		fmt.Printf("%d->", n.value)
		n = n.next
	}
	fmt.Println("nil")
}

func TestReorderList(t *testing.T) {
	ReorderList(makeList(31))
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
		} else {
			p.next = n
			p = n
		}
	}
	return h
}

type node struct {
	value int
	next  *node
}
