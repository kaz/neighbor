package main

type (
	Index interface {
		Build(data []uint64)
		Lookup(value uint64, tolerance int) int
	}
)
