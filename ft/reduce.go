package ft

func Reduce[T any](f func(T, T) T, input []T, initial ...T) (result T) {
	values := append(initial, input...)
	if len(values) == 0 {
		return
	}

	acc := values[0]
	for _, v := range values[1:] {
		acc = f(acc, v)
	}
	return acc
}
