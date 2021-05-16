package main

import "testing"

func BenchmarkFizzBuzz(b *testing.B) {
	FizzBuzz(b.N)
}

func BenchmarkFizzBuzz_v2(b *testing.B) {
	FizzBuzz_v2(b.N)
}
