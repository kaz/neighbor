package main

type (
	V3_0 struct {
		data []uint64
	}
)

func (v *V3_0) Build(data []uint64) {
	v.data = data
}
func (v *V3_0) Lookup(value uint64, tolerance int) int {
	return LookupSIMD(v.data, value, tolerance)
}
