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
func (v *V2_1) Lookup(value uint64, tolerance int) []uint64 {
	return v.tree.Find1(value, tolerance)
}

func (t *Tree) Find1(value uint64, tolerance int) []uint64 {
	if t.root == nil {
		return []uint64{}
	}

	result := []uint64{}
	candidates := []*Node{t.root}

	for len(candidates) > 0 {
		n := candidates[0]
		candidates = candidates[1:]

		distance := bits.OnesCount64(n.value ^ value)
		if distance <= tolerance {
			result = append(result, n.value)
		}

		for key, child := range n.children {
			if distance-tolerance <= key && key <= distance+tolerance {
				candidates = append(candidates, child)
			}
		}
	}

	return result
}
