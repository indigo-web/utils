package uf

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestB2S(t *testing.T) {
	bytes := []byte("Hello, world!")
	require.Equal(t, "Hello, world!", B2S(bytes))
}
