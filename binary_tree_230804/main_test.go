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
	if n != nil {
		visit(n)
		preOrder(n.left)
		preOrder(n.right)
	}
}

func preOrder2(n *node) {
	stack := NewStack()
	p := n
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			stack.Push(p)
			visit(p)
			p = p.left
		} else {
			x, err := stack.Pop()
			if err != nil {
				fmt.Print(err.Error())
				p = nil
				break
			}
			p = x.(*node)
			p = p.right
		}
	}
}

func inOrder(n *node) {
	if n != nil {
		inOrder(n.left)
		visit(n)
		inOrder(n.right)
	}
}

func inOrder2(n *node) {
	stack := NewStack()
	p := n
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			stack.Push(p)
			p = p.left
		} else {
			x, err := stack.Pop()
			if err != nil {
				fmt.Print(err.Error())
				p = nil
				break
			}
			p = x.(*node)
			visit(p)
			p = p.right
		}
	}
}

func postOrder(n *node) {
	if n != nil {
		postOrder(n.left)
		postOrder(n.right)
		visit(n)
	}
}

func postOrder2(n *node) {
	stack := NewStack()
	p := n
	var r *node = nil
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			stack.Push(p)
			p = p.left
		} else {
			x := stack.Top()
			if x == nil {
				break
			}
			p = x.(*node)
			if p.right != nil && p.right != r {
				p = p.right
				stack.Push(p)
				p = p.left
			} else {
				x, err := stack.Pop()
				if err != nil {
					fmt.Print(err.Error())
					p = nil
					break
				}
				p = x.(*node)
				visit(p)
				r = p
				p = nil
			}
		}
	}
}

func postOrder22(n *node) {
	stack := NewStack()
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
				x, _ := stack.Pop()
				p = x.(*node)
				visit(p)
				r = p
				p = nil
			}
		}
	}
}

func visit(n *node) {
	fmt.Printf("Visiting Node: %s\n", n.value)
}

func makeBinaryTree() *node {
	ln1 := &node{
		value: "D",
	}
	ln2 := &node{
		value: "E",
	}
	ln := &node{
		value: "B",
		left:  ln1,
		right: ln2,
	}
	rn := &node{
		value: "C",
	}
	n := &node{
		value: "A",
		left:  ln,
		right: rn,
	}
	return n
}

func TestPreOrder(t *testing.T) {
	n := makeBinaryTree()
	preOrder(n)
}

func TestPreOrder2(t *testing.T) {
	n := makeBinaryTree()
	preOrder2(n)
}

func TestInOrder(t *testing.T) {
	n := makeBinaryTree()
	inOrder(n)
}

func TestInOrder2(t *testing.T) {
	n := makeBinaryTree()
	inOrder2(n)
}

func TestPostOrder(t *testing.T) {
	n := makeBinaryTree()
	postOrder(n)
}

func TestPostOrder2(t *testing.T) {
	n := makeBinaryTree()
	postOrder2(n)
}

func TestPostOrder22(t *testing.T) {
	n := makeBinaryTree()
	postOrder22(n)
}
