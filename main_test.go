package main

import (
	"testing"
)

func BenchmarkCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func BenchmarkCase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
