package main

import (
	"testing"
)

func TestV1_0(t *testing.T) {
	genTest(t, &V1_0{})
}

func BenchmarkV1_0(b *testing.B) {
	genBenchmark(b, &V1_0{})
}
