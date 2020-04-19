package porter2

import (
	"fmt"
	"testing"
)

var BenchIntResult int

func BenchmarkFindRException(b *testing.B) {
	for idx, tc := range []struct {
		in []byte
	}{
		{in: []byte("fooba")},
		{in: []byte("foobar")},
		{in: []byte("gooba")},
		{in: []byte("goobar")},
		{in: []byte("communism")},
	} {
		b.Run(fmt.Sprintf("%d/%s", idx, tc.in), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BenchIntResult = findRException(tc.in)
			}
		})
	}
}
