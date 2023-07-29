package arena

// Arena is simply a big slice static-sized slice. It behaves just like built-in append()
// function, except encapsulating its internal implementation, so no slice is being returned
// to user - just boolean flag whether newly appended data exceeds size limits.
type Arena[T any] struct {
	memory     []T
	begin, pos int

	maxSize int
}

// NewArena returns a new arena instance
func NewArena[T any](initialSpace, maxSpace int) *Arena[T] {
	return &Arena[T]{
		memory:  make([]T, initialSpace),
		maxSize: maxSpace,
	}
}

// Append appends bytes to a buffer. In case of exceeding the maximal size, false is returned
// and data isn't written
func (a *Arena[T]) Append(elements ...T) (ok bool) {
	if a.pos+len(elements) > len(a.memory) {
		if len(a.memory)+len(elements) >= a.maxSize {
			return false
		}

		copy(a.memory[a.pos:], elements)
		// rely on the built-in slice growth mechanism. This sometimes exceeds the size limit
		// without notifying the user anyhow, but this doesn't bother anybody at the moment
		a.memory = append(a.memory, elements[len(a.memory)-a.pos:]...)
		a.pos += len(elements)

		return true
	}

	copy(a.memory[a.pos:], elements)
	a.pos += len(elements)

	return true
}

// SegmentLength returns a number of bytes, taken by current segment, calculated as a difference
// between the beginning of the current segment and the current pointer
func (a *Arena[T]) SegmentLength() int {
	return a.pos - a.begin
}

// Finish completes current segment, returning its value
func (a *Arena[T]) Finish() []T {
	segment := a.memory[a.begin:a.pos]
	a.begin = a.pos

	return segment
}

// Clear just resets the pointers, so old values may be overridden by new ones.
func (a *Arena[T]) Clear() {
	a.begin = 0
	a.pos = 0
}
