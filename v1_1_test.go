package main

import (
	"testing"
)

func TestV1_1(t *testing.T) {
	genTest(t, &V1_0{})
}

func BenchmarkV1_1(b *testing.B) {
	genBenchmark(b, &V1_0{})
}
