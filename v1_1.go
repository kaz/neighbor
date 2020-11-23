package main

type (
	V1_1 struct {
		data []uint64
	}
)

func (v *V1_1) Build(data []uint64) {
	v.data = data
}
func (v *V1_1) Lookup(value uint64, tolerance int) int {
	return LookupPopcnt(v.data, value, tolerance)
}
