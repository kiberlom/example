package main

import "testing"

//go test -bench=. .
//go test -bench=. -benchmem .

func BenchmarkSilo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Silo()
	}
}

func BenchmarkS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Filo()
	}
}
