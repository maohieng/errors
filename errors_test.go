package errs

import (
	"errors"
	"testing"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const ops Op = "Oop"
		New(errors.New("Test error"), ops, KindNotAllowed)
	}
}

func BenchmarkE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const ops Op = "Oop"
		E(errors.New("Test error"), ops, KindNotAllowed)
	}
}
