package main

import "math/bits"

type (
	V2_0 struct {
		tree *Tree
	}

	Tree struct {
		root *Node
	}
	Node struct {
		value    uint64
		children map[int]*Node
	}
)

func (v *V2_0) Build(data []uint64) {
	v.tree = &Tree{}
	for _, ent := range data {
		v.tree.Add(ent)
	}
}
func (v *V2_0) Lookup(value uint64, tolerance int) int {
	return v.tree.Find(value, tolerance)
}

func NewNode(value uint64) *Node {
	return &Node{
		value:    value,
		children: map[int]*Node{},
	}
}

func (t *Tree) Add(value uint64) {
	if t.root == nil {
		t.root = NewNode(value)
		return
	}
	t.root.Add(value)
}
func (n *Node) Add(value uint64) {
	distance := bits.OnesCount64(n.value ^ value)
	child, ok := n.children[distance]
	if !ok {
		n.children[distance] = NewNode(value)
		return
	}
	child.Add(value)
}

func (t *Tree) Find(value uint64, tolerance int) int {
	if t.root == nil {
		return 0
	}
	return t.root.Find(value, tolerance)
}
func (n *Node) Find(value uint64, tolerance int) int {
	distance := bits.OnesCount64(n.value ^ value)

	result := 0
	if distance <= tolerance {
		result++
	}

	for key, child := range n.children {
		if distance-tolerance <= key && key <= distance+tolerance {
			result += child.Find(value, tolerance)
		}
	}

	return result
}
