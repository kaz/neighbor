package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var (
	testsLen    = 100
	databaseLen = 1_000_000
	benchLen    = 10_000_000

	tests    []uint64
	database []uint64
	bench    []uint64

	reference = &V1_0{}

	tolerances = []int{0, 1, 2, 4, 8, 16}
)

func TestMain(m *testing.M) {
	tests = createDataSet(testsLen)
	database = createDataSet(databaseLen)
	bench = createDataSet(benchLen)

	reference.Build(database)

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

func genTest(t *testing.T, idx Index) {
	idx.Build(database)
	for _, tolerance := range tolerances {
		t.Run(fmt.Sprintf("tolerance=%d", tolerance), func(t *testing.T) {
			for _, ent := range tests {
				actual := idx.Lookup(ent, tolerance)
				expected := reference.Lookup(ent, tolerance)
				if actual != expected {
					t.Errorf("ent=%d, actual=%d, expected=%d", ent, actual, expected)
				}
			}
		})
	}
}

func genBenchmark(b *testing.B, idx Index) {
	idx.Build(database)
	for _, tolerance := range tolerances {
		b.Run(fmt.Sprintf("tolerance=%d", tolerance), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				idx.Lookup(bench[i], tolerance)
			}
		})
	}
}
