package main

import "math/bits"

type (
	V1_0 struct {
		data []uint64
	}
)

func (v *V1_0) Build(data []uint64) {
	v.data = data
}
func (v *V1_0) Lookup(value uint64, tolerance int) int {
	result := 0
	for _, ent := range v.data {
		if bits.OnesCount64(ent^value) <= tolerance {
			result++
		}
	}
	return result
}
