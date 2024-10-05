package buffer

// Buffer is simply a big slice, that can grow up to some limit, specified via
// constructor. Mainly used for cases, when we want to store multiple data pieces
// in the buffer, without knowing the actual size of each. So basically, it
// operates completely as append does, but encapsulates some important logic, like
// overflow checks and markers
type Buffer struct {
	memory  []byte
	begin   int
	maxSize int
}

func New(initialSpace, maxSpace int) *Buffer {
	return &Buffer{
		memory:  make([]byte, 0, initialSpace),
		maxSize: maxSpace,
	}
}

// Append appends bytes to a buffer. In case of exceeding the maximal size, false is returned
// and data isn't written
func (a *Buffer) Append(elements []byte) (ok bool) {
	if len(a.memory)+len(elements) > a.maxSize {
		return false
	}

	a.memory = append(a.memory, elements...)
	return true
}

// SegmentLength returns a number of bytes, taken by current segment, calculated as a difference
// between the beginning of the current segment and the current pointer
func (a *Buffer) SegmentLength() int {
	return len(a.memory) - a.begin
}

// Trunc truncates the last n bytes from the current segment, guaranting that data of previous segments stays intact
func (a *Buffer) Trunc(n int) {
	if seglen := a.SegmentLength(); n > seglen {
		n = seglen
	}

	a.memory = a.memory[:len(a.memory)-n]
}

// Discard discards current segment, and brings begin mark back by n bytes
func (a *Buffer) Discard(n int) {
	if n > a.begin {
		n = a.begin
	}

	a.begin -= n
	a.memory = a.memory[:a.begin]
}

// Preview returns current segment without moving the head
func (a *Buffer) Preview() []byte {
	return a.memory[a.begin:]
}

// Finish completes current segment, returning its value
func (a *Buffer) Finish() []byte {
	segment := a.memory[a.begin:]
	a.begin = len(a.memory)

	return segment
}

// Clear just resets the pointers, so old values may be overridden by new ones.
func (a *Buffer) Clear() {
	a.begin = 0
	a.memory = a.memory[:0]
}
