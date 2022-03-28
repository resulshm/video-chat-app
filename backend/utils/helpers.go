package utils

// TODO: improve email validation!
func IsEmailValid(email string) bool {

	var (
		hasMinlen = false
		hasMail   = false
		hasDot    = false
	)

	if len(email) > 8 {
		hasMinlen = true
	}

	for _, char := range email {
		switch char {
		case '@':
			hasMail = true
		case '.':
			hasDot = true
		}
	}

	return hasMinlen && hasMail && hasDot
}
