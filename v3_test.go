package main

import (
	"testing"
)

func TestV3(t *testing.T) {
	genTest(t, &V3{})
}

func BenchmarkV3(b *testing.B) {
	genBenchmark(b, &V3{})
}
