package main

import (
	"unsafe"
)

//go:noescape
func __find(haystack, haystackLen, needle, tolerance, result unsafe.Pointer)

func Find(haystack []uint64, needle uint64, tolerance int) int {
	var result int
	__find(
		unsafe.Pointer(&haystack[0]),
		unsafe.Pointer(uintptr(len(haystack))),
		unsafe.Pointer(uintptr(needle)),
		unsafe.Pointer(uintptr(tolerance)),
		unsafe.Pointer(&result),
	)
	return result
}
