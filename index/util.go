package index

func keyIsInvalid(key []byte) bool {
	if key == nil || len(key) == 0 {
		return true
	}
	return false
}
