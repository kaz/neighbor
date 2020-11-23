package main

import "math/bits"

type (
	V1 struct {
		data []uint64
	}
)

func (v *V1) Build(data []uint64) {
	v.data = data
}
func (v *V1) Lookup(value uint64, tolerance int) []uint64 {
	result := []uint64{}
	for _, ent := range v.data {
		if bits.OnesCount64(ent^value) <= tolerance {
			result = append(result, ent)
		}
	}
	return result
}
