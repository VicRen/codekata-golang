package main

import (
	"fmt"
	"testing"
)

func moveList(k int, head *node) {
	printList(head)
	if head == nil {
		return
	}
	cnt := head.Len()
	if k == cnt || k == 0 {
		return
	}

	p := head
	var r *node
	if k > cnt {
		k = k % cnt
	}
	for i := 0; i < k; i++ {
		r = p
		p = p.next
	}
	r.next = nil
	printList(p)
	printList(head)
	r = head
	head = p
	for p != nil && p.next != nil {
		p = p.next
	}
	p.next = r
	printList(head)
}

func TestMoveList(t *testing.T) {
	moveList(4, makeList1())
	moveList(0, makeList1())
	moveList(2, makeList2())
	moveList(5, makeList2())
}

func printList(head *node) {
	for head != nil {
		fmt.Printf("%d->", head.value)
		head = head.next
	}
	fmt.Println("nil")
}

func makeList2() *node {
	return &node{
		next: &node{
			next: &node{
				next: &node{
					next: &node{
						value: 4,
					},
					value: 3,
				},
				value: 2,
			},
			value: 1,
		},
		value: 0,
	}
}

func makeList1() *node {
	return &node{
		next: &node{
			next: &node{
				value: 2,
			},
			value: 1,
		},
		value: 0,
	}
}

type node struct {
	value int
	next  *node
}

func (n *node) Len() int {
	ret := 0
	p := n
	for p != nil {
		p = p.next
		ret++
	}
	return ret
}
