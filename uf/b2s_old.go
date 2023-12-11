//go:build !go1.20

package uf

import "unsafe"

func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
