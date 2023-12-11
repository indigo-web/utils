package buffer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func pushSegment(t *testing.T, buff *Buffer[byte], text string) {
	ok := buff.Append([]byte(text)...)
	require.True(t, ok)
	segment := buff.Finish()
	require.Equal(t, text, string(segment))
}

func TestBuffer(t *testing.T) {
	t.Run("NoOverflow", func(t *testing.T) {
		buff := New[byte](10, 20)
		pushSegment(t, buff, "Hello")
		pushSegment(t, buff, "Here")
	})

	t.Run("YesOverflow", func(t *testing.T) {
		buff := New[byte](10, 20)
		// "Hello, World!" is 13 characters length, so it will force the Buffer
		// to grow an underlying slice
		pushSegment(t, buff, "Hello, ")
		pushSegment(t, buff, "World!")
	})

	t.Run("SizeLimitOverflow", func(t *testing.T) {
		buff := New[byte](10, 20)
		pushSegment(t, buff, "Hello, ")
		pushSegment(t, buff, "World!")
		pushSegment(t, buff, "Lorem ")
		// at this point, we have reached 19 elements in underlying slice
		ok := buff.Append([]byte("overflow")...)
		require.False(t, ok)
	})

	t.Run("SegmentLength", func(t *testing.T) {
		buff := New[byte](10, 20)
		require.True(t, buff.Append([]byte("Hello, ")...))
		require.True(t, buff.Append([]byte("World!")...))
		require.Equal(t, 13, buff.SegmentLength())
	})

	t.Run("Discard", func(t *testing.T) {
		testDiscard(t, 13)
	})

	t.Run("BigDiscard", func(t *testing.T) {
		testDiscard(t, 50)
	})
}

func testDiscard(t *testing.T, n int) {
	buff := New[byte](10, 20)
	require.True(t, buff.Append([]byte("Hello, world!")...))
	segment := buff.Finish()
	buff.Discard(n)
	require.True(t, buff.Append([]byte("Hello!")...))
	newSegment := buff.Finish()
	require.Equal(t, "Hello!", string(newSegment))
	require.Equal(t, "Hello! world!", string(segment))
}
