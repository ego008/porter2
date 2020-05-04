package porter2

import (
	"fmt"
	"testing"
)

func TestRemoveAposEtc(t *testing.T) {
	for idx, tc := range []struct {
		in  string
		out string
	}{
		{"foo", "foo"},
		{"foos", "foos"},
		{"foo's", "foo"},
		{"foos's", "foos"},
		{"'s", ""},
		{"s's", "s"},
		{"'s'", ""},
		{"'", ""},
		{"s'", "s"},
		{"s's'", "s"},
	} {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			r1 := removeSuffix_apos_s_apos_all([]byte(tc.in))
			r2 := removeSuffix_apos_s_apos([]byte(tc.in))
			r2 = removeSuffix_apos_s(r2)
			r2 = removeSuffix_apos(r2)
			if string(r1) != tc.out {
				t.Fatal(string(r1), "!=", tc.out)
			}
			if string(r1) != string(r2) {
				t.Fatal(string(r1), "!=", string(r2))
			}
		})
	}
}

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
