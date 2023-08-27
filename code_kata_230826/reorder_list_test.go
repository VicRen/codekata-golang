package main

import (
	"fmt"
	"testing"
)

func reorderList(head *node) {
	if head == nil || head.next == nil || head.next.next == nil {
		return
	}

	f, s := head, head
	for f.next != nil && f.next.next != nil {
		f = f.next.next
		s = s.next
	}

	sh := s.next
	printList(sh)
	s.next = nil
	printList(head)
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

func TestReorderList(t *testing.T) {
	printList(makeNode(10))
	reorderList(makeNode(10))
	reorderList2(makeNode(21))
}

func printList(head *node) {
	for head != nil {
		fmt.Printf("%d->", head.value)
		head = head.next
	}
	fmt.Println("nil")
}

func reorderList2(head *node) {
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
	for p2.next != nil {
		next1, next2 := p1.next, p2.next
		p1.next = p2
		p2.next = next1
		p1, p2 = next1, next2
	}
	printList(head)
}

func makeNode(count int) *node {
	var h *node = nil
	var p *node = nil
	for i := 0; i < count; i++ {
		n := &node{
			value: i,
		}
		if h == nil {
			h = n
			p = h
		} else {
			p.next = n
			p = p.next
		}
	}
	return h
}

type node struct {
	value int
	next  *node
}
