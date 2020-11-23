package main

import (
	"testing"
)

func TestV2(t *testing.T) {
	genTest(t, &V2{})
}

func BenchmarkV2(b *testing.B) {
	genBenchmark(b, &V2{})
}
