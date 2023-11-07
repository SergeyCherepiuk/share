package ot

import "testing"

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinimumEditDistance([]byte("kitten"), []byte("sitting"))
	}
}
