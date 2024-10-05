package buffer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func pushSegment(t *testing.T, buff *Buffer, text string) {
	ok := buff.Append([]byte(text))
	require.True(t, ok)
	segment := buff.Finish()
	require.Equal(t, text, string(segment))
}

func TestBuffer(t *testing.T) {
	t.Run("no overflow", func(t *testing.T) {
		buff := New(10, 20)
		pushSegment(t, buff, "Hello")
		pushSegment(t, buff, "Here")
	})

	t.Run("with overflow", func(t *testing.T) {
		buff := New(10, 20)
		// "Hello, World!" is 13 characters length, so it will force the Buffer
		// to grow an underlying slice
		pushSegment(t, buff, "Hello, ")
		pushSegment(t, buff, "World!")
	})

	t.Run("overflow over the limit", func(t *testing.T) {
		buff := New(10, 20)
		pushSegment(t, buff, "Hello, ")
		pushSegment(t, buff, "World!")
		pushSegment(t, buff, "Lorem ")
		// at this point, we have reached 19 elements in underlying slice
		ok := buff.Append([]byte("overflow"))
		require.False(t, ok)
	})

	t.Run("segment length", func(t *testing.T) {
		buff := New(10, 20)
		require.True(t, buff.Append([]byte("Hello, ")))
		require.True(t, buff.Append([]byte("World!")))
		require.Equal(t, 13, buff.SegmentLength())
	})

	t.Run("discard segment", func(t *testing.T) {
		testDiscard(t, 13)
		testDiscard(t, 50)
	})

	t.Run("truncate", func(t *testing.T) {
		testTrunc(t, 1)
		testTrunc(t, 5)
	})
}

func testDiscard(t *testing.T, n int) {
	buff := New(10, 20)
	require.True(t, buff.Append([]byte("Hello, world!")))
	segment := buff.Finish()
	buff.Discard(n)
	require.True(t, buff.Append([]byte("Hello!")))
	newSegment := buff.Finish()
	require.Equal(t, "Hello!", string(newSegment))
	require.Equal(t, "Hello! world!", string(segment))
}

func testTrunc(t *testing.T, n int) {
	buff := New(10, 20)
	require.True(t, buff.Append([]byte("Hello, world!")))
	segment := buff.Finish()
	require.True(t, buff.Append([]byte("Hi?")))
	buff.Trunc(n)
	require.Equal(t, "Hello, world!", string(segment))

	orig := "Hi?"
	if n > len(orig) {
		n = len(orig)
	}

	require.Equal(t, orig[:len(orig)-n], string(buff.Finish()))
}
