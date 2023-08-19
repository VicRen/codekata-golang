package main

import "testing"

func preOrder(n *node) string {
	stack := MakeStack()
	p := n
	var ret string
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			ret += p.value
			stack.Push(p)
			p = p.left
		} else {
			p = stack.Pop().(*node)
			p = p.right
		}
	}
	return ret
}

func inOrder(n *node) string {
	stack := MakeStack()
	p := n
	var ret string
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			stack.Push(p)
			p = p.left
		} else {
			p = stack.Pop().(*node)
			ret += p.value
			p = p.right
		}
	}
	return ret
}

func postOrder(n *node) string {
	stack := MakeStack()
	p := n
	var ret string
	var r *node
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
				ret += p.value
				r = p
				p = nil
			}
		}
	}
	return ret
}

func TestPreOrder(t *testing.T) {
	want := "ABDECFG"
	got := preOrder(makeBinaryTree())
	if got != want {
		t.Errorf("preOrder = %s, want %s\n", got, want)
	}
}

func TestInOrder(t *testing.T) {
	want := "DBEAFCG"
	got := inOrder(makeBinaryTree())
	if got != want {
		t.Errorf("inOrder = %s, want %s\n", got, want)
	}
}

func TestPostOrder(t *testing.T) {
	want := "DEBFGCA"
	got := postOrder(makeBinaryTree())
	if got != want {
		t.Errorf("postOrder = %s, want %s\n", got, want)
	}
}

type node struct {
	left  *node
	right *node
	value string
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
