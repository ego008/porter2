package porter2

var caseFold = [256]byte{}

func init() {
	for i := 0; i < 256; i++ {
		b := byte(i)
		if b >= 'A' && b <= 'Z' {
			b += 'a' - 'A'
		}
		caseFold[i] = b
	}
}

// toLower performs an ASCII-only lowercase conversion on s.
func toLower(s []byte) []byte {
	for i, b := range s {
		s[i] = caseFold[b]
	}
	return s
}
