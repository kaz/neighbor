package main

import "math/bits"

func v1(data []uint64, tolerance int) int {
	result := 0
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			if bits.OnesCount64(data[i]^data[j]) <= tolerance {
				result += 1
			}
		}
	}
	return result
}
