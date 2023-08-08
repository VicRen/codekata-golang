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

func PreOrder(n *node) {
	if n == nil {
		return
	}
	visit(n)
	PreOrder(n.left)
	PreOrder(n.right)
}

func PreOrder2(n *node) {
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

func PreOrder22(n *node) {
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

func InOrder(n *node) {
	if n == nil {
		return
	}
	InOrder(n.left)
	visit(n)
	InOrder(n.right)
}

func InOrder2(n *node) {
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

func InOrder22(n *node) {
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

func PostOrder(n *node) {
	if n == nil {
		return
	}
	PostOrder(n.left)
	PostOrder(n.right)
	visit(n)
}

func PostOrder2(n *node) {
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

func PostOrder22(n *node) {
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

func postOrder222(n *node) {
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

func postOrder2222(n *node) {
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
	PreOrder(makeBinaryTree())
}

func TestPreOrder2(t *testing.T) {
	PreOrder2(makeBinaryTree())
}

func TestPreOrder22(t *testing.T) {
	PreOrder22(makeBinaryTree())
}

func TestInOrder(t *testing.T) {
	InOrder(makeBinaryTree())
}

func TestInOrder2(t *testing.T) {
	InOrder2(makeBinaryTree())
}

func TestInOrder22(t *testing.T) {
	InOrder22(makeBinaryTree())
}

func TestPostOrder(t *testing.T) {
	PostOrder(makeBinaryTree())
}

func TestPostOrder2(t *testing.T) {
	PostOrder2(makeBinaryTree())
}

func TestPostOrder22(t *testing.T) {
	PostOrder22(makeBinaryTree())
}

func TestPostOrder222(t *testing.T) {
	postOrder222(makeBinaryTree())
}

func TestPostOrder2222(t *testing.T) {
	postOrder2222(makeBinaryTree())
}

func visit(n *node) {
	fmt.Printf("visiting node: %s\n", n.value)
}

func makeBinaryTree() *node {
	return &node{
		left: &node{
			left: &node{
				value: "D",
			},
			value: "B",
		},
		right: &node{
			right: &node{
				value: "E",
			},
			value: "C",
		},
		value: "A",
	}
}
