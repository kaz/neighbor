package main

import (
	"testing"
)

func TestV2_0(t *testing.T) {
	genTest(t, &V2_0{})
}

func BenchmarkV2_0(b *testing.B) {
	genBenchmark(b, &V2_0{})
}
