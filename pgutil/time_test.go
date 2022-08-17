package pgutil_test

import (
	"github.com/go-season/pg/v3/pgutil"
	"testing"
)

func BenchmarkParseTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pgutil.ParseTime([]byte("2001-02-03 04:05:06+07"))
	}
}
