package main

import "math/bits"

type (
	V2_1 struct {
		tree *Tree
	}
)

func (v *V2_1) Build(data []uint64) {
	v.tree = &Tree{}
	for _, ent := range data {
		v.tree.Add(ent)
	}
}
func (v *V2_1) Lookup(value uint64, tolerance int) int {
	return v.tree.FindLoop(value, tolerance)
}

func (t *Tree) FindLoop(value uint64, tolerance int) int {
	if t.root == nil {
		return 0
	}

	result := 0
	candidates := []*Node{t.root}

	for len(candidates) > 0 {
		n := candidates[0]
		candidates = candidates[1:]

		distance := bits.OnesCount64(n.value ^ value)
		if distance <= tolerance {
			result++
		}

		for key, child := range n.children {
			if distance-tolerance <= key && key <= distance+tolerance {
				candidates = append(candidates, child)
			}
		}
	}

	return result
}
