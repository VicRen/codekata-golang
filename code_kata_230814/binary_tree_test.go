package main

import (
	"fmt"
	"testing"
)

func preOrder(n *node) {
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
	preOrder(makeBinaryTree())
}

func TestInOrder(t *testing.T) {
	inOrder(makeBinaryTree())
}

func TestPostOrder(t *testing.T) {
	postOrder(makeBinaryTree())
}

func visit(n *node) {
	if n == nil {
		return
	}
	fmt.Printf("visiting node %s\n", n.value)
}

func makeBinaryTree() *node {
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

type node struct {
	value string
	left  *node
	right *node
}
