package benchmark

import "testing"

func BenchmarkMakeSliceWithAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MakeSliceWithAlloc()
	}
}

func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MakeSliceWithoutAlloc()
	}
}
