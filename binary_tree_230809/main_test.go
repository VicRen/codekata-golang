package main

import (
	"fmt"
	"testing"
)

type node struct {
	left  *node
	right *node
	value string
}

func preOrder(n *node) {
	if n == nil {
		return
	}
	visit(n)
	preOrder(n.left)
	preOrder(n.right)
}

func preOrder2(n *node) {
	stack := MakeStack()
	p := n
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			visit(p)
			stack.Push(p)
			p = p.left
		} else {
			p = stack.Pop().(*node)
			p = p.right
		}
	}
}

func inOrder(n *node) {
	if n == nil {
		return
	}
	inOrder(n.left)
	visit(n)
	inOrder(n.right)
}

func inOrder2(n *node) {
	stack := MakeStack()
	p := n
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			stack.Push(p)
			p = p.left
		} else {
			p = stack.Pop().(*node)
			visit(p)
			p = p.right
		}
	}
}

func postOrder(n *node) {
	if n == nil {
		return
	}
	postOrder(n.left)
	postOrder(n.right)
	visit(n)
}

func postOrder2(n *node) {
	stack := MakeStack()
	p := n
	var r *node = nil
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			stack.Push(p)
			p = p.left
		} else {
			p = stack.Top().(*node)
			if p.right != nil && p.right != r {
				p = p.right
				stack.Push(p)
				p = p.left
			} else {
				p = stack.Pop().(*node)
				visit(p)
				r = p
				p = nil
			}
		}
	}
}

func TestPreOrder(t *testing.T) {
	preOrder(makeBinary())
}

func TestPreOrder2(t *testing.T) {
	preOrder2(makeBinary())
}

func TestInOrder(t *testing.T) {
	inOrder(makeBinary())
}

func TestInOrder2(t *testing.T) {
	inOrder2(makeBinary())
}

func TestPostOrder(t *testing.T) {
	postOrder(makeBinary())
}

func TestPostOrder2(t *testing.T) {
	postOrder2(makeBinary())
}

func visit(n *node) {
	if n == nil {
		return
	}
	fmt.Printf("visiting node %s\n", n.value)
}

func makeBinary() *node {
	return &node{
		left: &node{
			left: &node{
				value: "D",
			},
			right: &node{
				value: "E",
			},
			value: "B",
		},
		right: &node{
			left: &node{
				value: "F",
			},
			right: &node{
				value: "G",
			},
			value: "C",
		},
		value: "A",
	}
}
