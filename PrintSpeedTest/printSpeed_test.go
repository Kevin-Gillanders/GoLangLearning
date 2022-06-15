package main

import "testing"

//Not suitable for benchmarking
func BenchmarkFmtPrint(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmtPrint(1000000)
	}

}
