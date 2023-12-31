package mapconv

// Keys is currently used only in http/decode/decoder.go, but
// this or other functions that may be added later may be used somewhere
// else
func Keys[K comparable, V any](from map[K]V) []K {
	keys := make([]K, 0, len(from))
	for key := range from {
		keys = append(keys, key)
	}

	return keys
}

func Clone[K comparable, V any](m map[K]V) map[K]V {
	newMap := make(map[K]V, len(m))

	for k, v := range m {
		newMap[k] = v
	}

	return newMap
}

func Copy[K comparable, V any](dst, src map[K]V) map[K]V {
	if dst == nil {
		dst = make(map[K]V, len(src))
	}

	for key, value := range src {
		dst[key] = value
	}

	return dst
}
