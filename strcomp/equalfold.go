package strcomp

func EqualFold(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	var i int

	for ; i < len(a)-8; i += 8 {
		if a[i]|0x20 != b[i]|0x20 {
			return false
		}
		if a[i+1]|0x20 != b[i+1]|0x20 {
			return false
		}
		if a[i+2]|0x20 != b[i+2]|0x20 {
			return false
		}
		if a[i+3]|0x20 != b[i+3]|0x20 {
			return false
		}
		if a[i+4]|0x20 != b[i+4]|0x20 {
			return false
		}
		if a[i+5]|0x20 != b[i+5]|0x20 {
			return false
		}
		if a[i+6]|0x20 != b[i+6]|0x20 {
			return false
		}
		if a[i+7]|0x20 != b[i+7]|0x20 {
			return false
		}
	}

	for ; i < len(a); i++ {
		if a[i]|0x20 != b[i]|0x20 {
			return false
		}
	}

	return true
}
