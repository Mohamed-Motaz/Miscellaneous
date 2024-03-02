package main

import "testing"

func BenchmarkOnTheHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnTheHeap()
	}
}

func BenchmarkOnTheStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnTheStack()
	}
}
