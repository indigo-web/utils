package arena

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func pushSegment(t *testing.T, arena *Arena, text string) {
	ok := arena.Append([]byte(text))
	require.True(t, ok)
	segment := arena.Finish()
	require.Equal(t, text, string(segment))
}

func TestArena(t *testing.T) {
	t.Run("NoOverflow", func(t *testing.T) {
		arena := NewArena(10, 20)
		pushSegment(t, arena, "Hello")
		pushSegment(t, arena, "Here")
	})

	t.Run("YesOverflow", func(t *testing.T) {
		arena := NewArena(10, 20)
		// "Hello, World!" is 13 characters length, so it will force the Arena
		// to grow an underlying slice
		pushSegment(t, arena, "Hello, ")
		pushSegment(t, arena, "World!")
	})

	t.Run("SizeLimitOverflow", func(t *testing.T) {
		arena := NewArena(10, 20)
		pushSegment(t, arena, "Hello, ")
		pushSegment(t, arena, "World!")
		pushSegment(t, arena, "Lorem ")
		// at this point, we have reached 19 elements in underlying slice
		ok := arena.Append([]byte("overflow"))
		require.False(t, ok)
	})
}
