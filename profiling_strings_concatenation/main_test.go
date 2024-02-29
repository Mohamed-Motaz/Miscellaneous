package main

import "testing"

func BenchmarkConcatStringsPlusEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringsPlusEqual("hello")
	}
}

func BenchmarkConcatStringsPlusEqualImproved(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringsPlusEqualImproved("hello")
	}
}

func BenchmarkConcatStringsBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringsBytesBuffer("hello")
	}
}

func BenchmarkConcatStringsBytesBufferImproved(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringsBytesBufferImproved("hello")
	}
}
