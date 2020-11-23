package main

type (
	V3 struct {
		data []uint64
	}
)

func (v *V3) Build(data []uint64) {
	v.data = data
}
func (v *V3) Lookup(value uint64, tolerance int) []uint64 {
	Find(v.data, value, tolerance)
	return []uint64{}
}
