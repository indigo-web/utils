package unreader

// Unreader allows to store extra bytes, and return them on the next call.
//
// It's ready to work right after the instantiation.
type Unreader struct {
	Pending []byte
}

// PendingOr returns pending data if available, otherwise returns the result of the passed function
func (u *Unreader) PendingOr(otherwise func() ([]byte, error)) (data []byte, err error) {
	if len(u.Pending) > 0 {
		data, u.Pending = u.Pending, nil
		return data, nil
	}

	return otherwise()
}

// Unread saves extra, that'll be returned on the next call
func (u *Unreader) Unread(b []byte) {
	u.Pending = b
}

// Reset discards all the pending data
func (u *Unreader) Reset() {
	u.Pending = nil
}
