package main

import "testing"

func BenchmarkPs1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ps1()
	}
}
