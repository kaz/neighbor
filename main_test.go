package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var (
	reference = v1

	tolerances = []int{8, 16, 32}
	lengths    = []int{2_000, 10_000, 50_000}

	dataSets = map[int][]uint64{}
	expected = map[int]map[int]int{}
)

func TestMain(m *testing.M) {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for _, len := range lengths {
		dataSets[len] = createDataSet(len)
		expected[len] = map[int]int{}
		for _, tolerance := range tolerances {
			wg.Add(1)
			go func(len, tolerance int) {
				mu.Lock()
				expected[len][tolerance] = reference(dataSets[len], tolerance)
				mu.Unlock()
				wg.Done()
			}(len, tolerance)
		}
	}

	wg.Wait()
	m.Run()
}

func createDataSet(len int) []uint64 {
	rand.Seed(time.Now().UnixNano())

	data := make([]uint64, len)
	for i := 0; i < len; i++ {
		data[i] = rand.Uint64()
	}
	return data
}

func genTest(t *testing.T, fn func([]uint64, int) int) {
	for len, data := range dataSets {
		for _, tolerance := range tolerances {
			t.Run(fmt.Sprintf("len=%d,tolerance=%d", len, tolerance), func(t *testing.T) {
				result := fn(data, tolerance)
				if result != expected[len][tolerance] {
					t.Errorf("actual=%d,expected=%d", result, expected[len][tolerance])
				}
			})
		}
	}
}

func genBenchmark(b *testing.B, fn func([]uint64, int) int) {
	for len, data := range dataSets {
		for _, tolerance := range tolerances {
			b.Run(fmt.Sprintf("len=%d,tolerance=%d", len, tolerance), func(b *testing.B) {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					fn(data, tolerance)
				}
			})
		}
	}
}
