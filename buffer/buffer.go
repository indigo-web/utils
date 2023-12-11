package buffer

// Buffer is simply a big slice, that can grow up to some limit, specified via
// constructor. Mainly used for cases, when we want to store multiple data pieces
// in the buffer, without knowing the actual size of each. So basically, it
// operates completely as append does, but encapsulates some important logic, like
// overflow checks and markers
type Buffer[T any] struct {
	memory     []T
	begin, pos int

	maxSize int
}

func New[T any](initialSpace, maxSpace int) *Buffer[T] {
	return &Buffer[T]{
		memory:  make([]T, initialSpace),
		maxSize: maxSpace,
	}
}

// Append appends bytes to a buffer. In case of exceeding the maximal size, false is returned
// and data isn't written
func (a *Buffer[T]) Append(elements ...T) (ok bool) {
	if a.pos+len(elements) > len(a.memory) {
		if a.pos+len(elements) > a.maxSize {
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
func (a *Buffer[T]) SegmentLength() int {
	return a.pos - a.begin
}

// Discard discards current segment, and returns begin mark back by n bytes
func (a *Buffer[T]) Discard(n int) {
	if n > a.begin {
		n = a.begin
	}

	a.begin -= n
	a.pos = a.begin
}

// Finish completes current segment, returning its value
func (a *Buffer[T]) Finish() []T {
	segment := a.memory[a.begin:a.pos]
	a.begin = a.pos

	return segment
}

// Clear just resets the pointers, so old values may be overridden by new ones.
func (a *Buffer[T]) Clear() {
	a.begin = 0
	a.pos = 0
}
