package arena

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func pushSegment(t *testing.T, arena *Arena[byte], text string) {
	ok := arena.Append([]byte(text)...)
	require.True(t, ok)
	segment := arena.Finish()
	require.Equal(t, text, string(segment))
}

func TestArena(t *testing.T) {
	t.Run("NoOverflow", func(t *testing.T) {
		arena := NewArena[byte](10, 20)
		pushSegment(t, arena, "Hello")
		pushSegment(t, arena, "Here")
	})

	t.Run("YesOverflow", func(t *testing.T) {
		arena := NewArena[byte](10, 20)
		// "Hello, World!" is 13 characters length, so it will force the Arena
		// to grow an underlying slice
		pushSegment(t, arena, "Hello, ")
		pushSegment(t, arena, "World!")
	})

	t.Run("SizeLimitOverflow", func(t *testing.T) {
		arena := NewArena[byte](10, 20)
		pushSegment(t, arena, "Hello, ")
		pushSegment(t, arena, "World!")
		pushSegment(t, arena, "Lorem ")
		// at this point, we have reached 19 elements in underlying slice
		ok := arena.Append([]byte("overflow")...)
		require.False(t, ok)
	})

	t.Run("SegmentLength", func(t *testing.T) {
		arena := NewArena[byte](10, 20)
		require.True(t, arena.Append([]byte("Hello, ")...))
		require.True(t, arena.Append([]byte("World!")...))
		require.Equal(t, 13, arena.SegmentLength())
	})

	t.Run("Discard", func(t *testing.T) {
		testDiscard(t, 13)
	})

	t.Run("BigDiscard", func(t *testing.T) {
		testDiscard(t, 50)
	})
}

func testDiscard(t *testing.T, n int) {
	arena := NewArena[byte](10, 20)
	require.True(t, arena.Append([]byte("Hello, world!")...))
	segment := arena.Finish()
	arena.Discard(n)
	require.True(t, arena.Append([]byte("Hello!")...))
	newSegment := arena.Finish()
	require.Equal(t, "Hello!", string(newSegment))
	require.Equal(t, "Hello! world!", string(segment))
}
