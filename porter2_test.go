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

func loadStemCases() (stemCases []StemCase) {
	filePairs := []struct {
		inFile, outFile string
	}{
		{"testdata/test_voc.txt", "testdata/test_output.txt"},
		{"testdata/test_accent_voc.txt", "testdata/test_accent_output.txt"},
		{"testdata/test_korean_voc.txt", "testdata/test_korean_output.txt"},
	}

	stemCases = make([]StemCase, 0, 10000)

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
			stemCases = append(stemCases, StemCase{a, b})
		}
	}

	// Additional cases:
	stemCases = append(stemCases, []StemCase{
		// UTF-8 shouldn't be mangled:
		strCase("naïve", "naïv"),
	}...)

	return stemCases
}

func TestStemBytes(t *testing.T) {
	// StemCases = []StemCase{{[]byte("dying"), []byte("die")}}

	for _, tc := range loadStemCases() {
		t.Run(fmt.Sprintf("%s/byte", tc.In), func(t *testing.T) {
			in := []byte(string(tc.In)) // Careful not to mutate the test case's memory!
			st := StemBytes(in, 0)
			if string(st) != string(tc.Out) {
				t.Errorf("\"%s\" expected %q got %q", string(tc.In), string(tc.Out), string(st))
			}
		})
	}
}

func TestStem(t *testing.T) {
	for _, tc := range loadStemCases() {
		t.Run(fmt.Sprintf("%s/str", tc.In), func(t *testing.T) {
			st := Stem(string(tc.In), 0)
			if string(st) != string(tc.Out) {
				t.Errorf("\"%s\" expected %q got %q", string(tc.In), string(tc.Out), string(st))
			}
		})
	}
}

var StemResult []byte
var StemStringResult string

func BenchmarkStemBytes(b *testing.B) {
	stemCases := loadStemCases()
	idx := 0
	sz := len(stemCases)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := stemCases[idx]
		StemResult = StemBytes(c.In, 0)
		idx++
		if idx >= sz {
			idx = 0
		}
	}
}

func BenchmarkStem(b *testing.B) {
	stemCases := loadStemCases()
	idx := 0
	sz := len(stemCases)

	var strCases = make([]string, sz)
	for idx, sc := range stemCases {
		strCases[idx] = string(sc.In)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := strCases[idx]
		StemStringResult = Stem(c, 0)
		idx++
		if idx >= sz {
			idx = 0
		}
	}
}
