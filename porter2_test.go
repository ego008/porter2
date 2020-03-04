// Copyright 2012 The Stemmer Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package porter2

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

type StemCase struct {
	In, Out []byte
}

func strCase(in, out string) StemCase {
	return StemCase{[]byte(in), []byte(out)}
}

var StemCases []StemCase

func init() {
	filePairs := []struct {
		inFile, outFile string
	}{
		{"testdata/test_voc.txt", "testdata/test_output.txt"},
		{"testdata/test_accent_voc.txt", "testdata/test_accent_output.txt"},
		{"testdata/test_korean_voc.txt", "testdata/test_korean_output.txt"},
	}

	StemCases = make([]StemCase, 0, 10000)

	for _, fp := range filePairs {
		voc, err := ioutil.ReadFile(fp.inFile)
		if err != nil {
			panic(err)
		}
		out, err := ioutil.ReadFile(fp.outFile)
		if err != nil {
			panic(err)
		}

		vocLines := bytes.Split(voc, []byte{'\n'})
		outLines := bytes.Split(out, []byte{'\n'})

		if len(vocLines) != len(outLines) {
			panic(nil)
		}

		for i := 0; i < len(vocLines); i++ {
			a, b := vocLines[i], outLines[i]
			StemCases = append(StemCases, StemCase{a, b})
		}
	}

	// Additional cases:
	StemCases = append(StemCases, []StemCase{
		// UTF-8 shouldn't be mangled:
		strCase("naïve", "naïv"),
	}...)
}

func TestStem(t *testing.T) {
	// StemCases = []StemCase{{[]byte("dying"), []byte("die")}}
	for _, tc := range StemCases {
		t.Run(fmt.Sprintf("%s", tc.In), func(t *testing.T) {
			st := Stem([]byte(tc.In), 0)
			if string(st) != string(tc.Out) {
				t.Errorf("\"%s\" expected %q got %q", string(tc.In), string(tc.Out), string(st))
			}
		})
	}
}

var StemResult []byte

func BenchmarkStem(b *testing.B) {
	idx := 0
	sz := len(StemCases)
	for i := 0; i < b.N; i++ {
		c := StemCases[idx]
		StemResult = Stem(c.In, 0)
		idx++
		if idx >= sz {
			idx = 0
		}
	}
}
