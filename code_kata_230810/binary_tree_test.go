package main

import (
	"fmt"
	"testing"
)

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

func levelOrder(n *node) {
	queue := MakeQueue()
	p := n
	queue.EnQueue(p)
	for !queue.IsEmpty() {
		p = queue.DeQueue().(*node)
		visit(p)
		if p.left != nil {
			queue.EnQueue(p.left)
		}
		if p.right != nil {
			queue.EnQueue(p.right)
		}
	}
}

func TestPreOrder(t *testing.T) {
	preOrder(makeBinaryTree())
}

func TestPreOrder2(t *testing.T) {
	preOrder2(makeBinaryTree())
}

func TestInOrder(t *testing.T) {
	inOrder(makeBinaryTree())
}

func TestInOrder2(t *testing.T) {
	inOrder2(makeBinaryTree())
}

func TestPostOrder(t *testing.T) {
	postOrder(makeBinaryTree())
}

func TestPostOrder2(t *testing.T) {
	postOrder2(makeBinaryTree())
}

func TestLevelOrder(t *testing.T) {
	levelOrder(makeBinaryTree())
}

type node struct {
	left  *node
	right *node
	value string
}

func visit(n *node) {
	if n == nil {
		panic("invalid node")
	}
	fmt.Printf("Visiting node %s\n", n.value)
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
