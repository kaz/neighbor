package main

import (
	"testing"
)

func TestV3_0(t *testing.T) {
	genTest(t, &V3_0{})
}

func BenchmarkV3_0(b *testing.B) {
	genBenchmark(b, &V3_0{})
}
