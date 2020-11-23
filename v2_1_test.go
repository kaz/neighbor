package main

import (
	"testing"
)

func TestV2_1(t *testing.T) {
	genTest(t, &V2_1{})
}

func BenchmarkV2_1(b *testing.B) {
	genBenchmark(b, &V2_1{})
}
