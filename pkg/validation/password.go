package pkg_validation

// ValidatePassword checks whether the provided password meets the minimum length
// and contains at least one uppercase letter, one lowercase letter, and one digit.
func ValidatePassword(password string, length int) bool {
	hasUpper := false
	hasLower := false
	hasNumber := false

	if len(password) < length {
		return false
	}

	// handmade password validator
	for _, c := range password {
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= '0' && c <= '9':
			hasNumber = true
		}
	}

	if !(hasUpper && hasLower && hasNumber) {
		return false
	}

	return true
}