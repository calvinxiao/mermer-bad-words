package main

import (
	"testing"

	"github.com/spaolacci/murmur3"
)

var _byte = []byte("https://www.aftership.com")

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = murmur3.Sum64(_byte)
	}
}

func BenchmarkCacheHasher(b *testing.B) {
	hasher := murmur3.New64()
	for i := 0; i < b.N; i++ {
		_ = hasher.Sum(_byte)
	}
}
