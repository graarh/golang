package trace

import (
	"testing"
	"fmt"
)

func BenchmarkBacktrace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		backtrace(0)
	}
}

func BenchmarkLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Error("aaa")
	}
}

func BenchmarkCommonError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Errorf("%s", "aaa")
	}
}
