package main

import "testing"

func BenchmarkSilo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Silo()
	}
}
