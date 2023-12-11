//go:build go1.20

package uf

import "unsafe"

func B2S(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
