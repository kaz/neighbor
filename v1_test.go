package main

import (
	"testing"
)

func TestV1(t *testing.T) {
	genTest(t, v1)
}

func BenchmarkV1(b *testing.B) {
	genBenchmark(b, v1)
}
