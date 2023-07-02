package ft

import "github.com/indigo-web/utils/constraint"

func accumulator[T constraint.Addable](prev T, curr T) T {
	return prev + curr
}

func Sum[T constraints.Addable](input []T) T {
	return Reduce(accumulator[T], input)
}
