package main

import "math/bits"

type (
	V2 struct {
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

func (v *V2) Build(data []uint64) {
	v.tree = &Tree{}
	for _, ent := range data {
		v.tree.Add(ent)
	}
}
func (v *V2) Lookup(value uint64, tolerance int) []uint64 {
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

func (t *Tree) Find(value uint64, tolerance int) []uint64 {
	if t.root == nil {
		return []uint64{}
	}
	return t.root.Find(value, tolerance)
}
func (n *Node) Find(value uint64, tolerance int) []uint64 {
	distance := bits.OnesCount64(n.value ^ value)

	result := []uint64{}
	if distance <= tolerance {
		result = append(result, n.value)
	}

	for key, child := range n.children {
		if distance-tolerance <= key && key <= distance+tolerance {
			result = append(result, child.Find(value, tolerance)...)
		}
	}

	return result
}
