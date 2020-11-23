package main

import (
	"testing"
)

func TestV4_0(t *testing.T) {
	genTest(t, &V4_0{})
}

func BenchmarkV4_0(b *testing.B) {
	genBenchmark(b, &V4_0{})
}
