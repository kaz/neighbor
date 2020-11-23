package main

import (
	"math/bits"
	"runtime"
	"sync"
	"sync/atomic"
)

type (
	V1_1 struct {
		data []uint64

		wg sync.WaitGroup
		mu sync.Mutex
	}
)

func (v *V1_1) Build(data []uint64) {
	v.data = data
}
func (v *V1_1) Lookup(value uint64, tolerance int) []uint64 {
	gptr := int64(-1)
	result := []uint64{}

	for i := 0; i < runtime.NumCPU(); i++ {
		v.wg.Add(1)
		go func() {
			defer v.wg.Done()
			for {
				ptr := atomic.AddInt64(&gptr, 1)
				if ptr >= int64(len(v.data)) {
					return
				}
				if bits.OnesCount64(v.data[ptr]^value) <= tolerance {
					v.mu.Lock()
					result = append(result, v.data[ptr])
					v.mu.Unlock()
				}
			}
		}()
	}

	v.wg.Wait()
	return result
}
