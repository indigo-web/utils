package buffer

import (
	"strings"
	"testing"
)

func BenchmarkBuffer(b *testing.B) {
	buff := New[byte](1024, 4096)
	smallString := []byte(strings.Repeat("a", 1023))
	bigString := []byte(strings.Repeat("a", 4095))

	b.Run("no overflow", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(smallString)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = buff.Append(smallString...)
			buff.Clear()
		}
	})

	b.Run("with overflow", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(bigString)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = buff.Append(bigString...)
			buff.Clear()
			buff.memory = buff.memory[0:1024:1024]
		}
	})
}
