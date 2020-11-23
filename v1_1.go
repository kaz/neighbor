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
func (v *V1_1) Lookup(value uint64, tolerance int) int {
	gptr := int64(-1)
	result := int64(0)

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
					atomic.AddInt64(&result, 1)
				}
			}
		}()
	}

	v.wg.Wait()
	return int(result)
}
