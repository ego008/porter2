package porter2

import (
	"fmt"
	"testing"
)

func BenchmarkSuffixPos_eedly(b *testing.B) {
	for idx, bc := range []struct {
		in []byte
	}{
		{in: []byte("wat")},
		{in: []byte("quack")},
		{in: []byte("indoodly")},
		{in: []byte("indee")},
		{in: []byte("indeed")},
		{in: []byte("indeedl")},
		{in: []byte("indeedly")},
	} {
		b.Run(fmt.Sprintf("%d", idx), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BenchIntResult = suffixPos_eedly(bc.in)
			}
		})
	}
}
