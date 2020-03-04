package porter2

import (
	"unsafe"
)

// StemString is the safe version of Stem. It operates on strings instead of bytes and
// copies the memory for you so you don't get crazy mutations you didn't expect or want.
//
// If you care about speed, you should use Stem instead and be careful with your memory.
func StemString(s string, flag StemFlag) string {
	out := Stem([]byte(s), flag)
	return *(*string)(unsafe.Pointer(&out))
}
