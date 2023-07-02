package ft

func Map[A any, B any](fun func(A) B, input []A) []B {
	result := make([]B, len(input))

	for i := range input {
		result[i] = fun(input[i])
	}

	return result
}

func Nop[T any](a T) T {
	return a
}
