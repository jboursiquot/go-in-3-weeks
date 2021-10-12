package fibonacci

import "testing"

func benchFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(i)
	}
}

func BenchmarkFib1(b *testing.B)  { benchFib(1, b) }
func BenchmarkFib10(b *testing.B) { benchFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchFib(20, b) }
