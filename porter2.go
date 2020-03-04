// Copyright 2012 The Stemmer Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package porter2 implements English (Porter2) stemmer, as described in
// http://snowball.tartarus.org/algorithms/english/stemmer.html
package porter2

import (
	"bytes"
	"unsafe"
)

type StemFlag int

const (
	UTF8Lower = 1 << iota
)

// Stem takes the string 'word' and stems it according to the porter2 rules.
//
// If you require raw speed and don't care about mutation, use StemBytes.
func Stem(s string, flag StemFlag) string {
	out := StemBytes([]byte(s), flag)
	return *(*string)(unsafe.Pointer(&out))
}

// StemBytes takes the byte slice 'word' and stems it according to the porter2 rules,
// then returns the slice truncated to the correct length.
//
// Warning: the byte slice passed is mutated! If you would prefer this not happen,
// clone the memory yourself before stemming:
//
//	out := Stem([]byte(string(word)), 0)
//
func StemBytes(word []byte, flag StemFlag) []byte {
	// XXX: ASCII-only seems OK to me, but using bytes.ToLower will potentially
	// reduce the size of the term space for non-ASCII terms, which do occur in
	// the data I'm using from time to time.
	var s []byte
	if flag&UTF8Lower != 0 {
		s = bytes.ToLower(word)
	} else {
		s = toLower(word)
	}

	// Is it exception?
	if rep, ex := exceptions1.Find(s); ex {
		return rep
	}
	if len(s) <= 2 {
		return word
	}
	if s[0] == '\'' {
		s = s[1:]
	}
	if s[0] == 'y' {
		s[0] = 'Y'
	}
	for i := 1; i < len(s); i++ {
		if isVowel[s[i-1]] && s[i] == 'y' {
			s[i] = 'Y'
		}
	}
	r1, r2 := getR1R2(s)

	// Step 0
	s = removeSuffix_apos_s_apos(s)
	s = removeSuffix_apos_s(s)
	s = removeSuffix_apos(s)

	// Step 1a
	if i := suffixPos_sses(s); i != -1 {
		// sses, replace by ss
		s = s[:i+2]
		goto step1b
	}
	{
		i := suffixPos_ied(s)
		if i == -1 {
			i = suffixPos_ies(s)
		}
		if i != -1 {
			// ied+   ies*
			// replace by i if preceded by more than one letter,
			// otherwise by ie (so ties -> tie, cries -> cri)
			if i > 1 {
				s = s[:i+1] // equivalent: append(s[:i], 'i')
			} else {
				s = s[:i+2] // equivalent: append(s[:i], 'i', 'e')
			}
			goto step1b
		}
	}
	if suffixPos_us(s) != -1 || suffixPos_ss(s) != -1 {
		// do nothing
		goto step1b
	}

	if i := suffixPos_s(s); i != -1 {
		if len(s) >= 3 && hasVowelBeforePos(s, len(s)-3) {
			s = s[:i]
		}
		goto step1b
	}

step1b:
	if isException2(s) {
		return s
	}

	// Step 1b
	if i := suffixPos_eed(s); i != -1 {
		if i >= r1 {
			s = append(s[:i], 'e', 'e')
		}
		goto step1c
	}
	if i := suffixPos_eedly(s); i != -1 {
		if i >= r1 {
			s = append(s[:i], 'e', 'e')
		}
		goto step1c
	}

	if found, idx, _ := step1bTree.Find(s); found {
		suf := step1bWords[idx]
		if len(s) > len(suf) && hasVowelBeforePos(s, len(s)-len(suf)-1) {
			s = s[:len(s)-len(suf)]
		} else {
			goto step1c
		}
		if suffixPos_at(s) != -1 || suffixPos_bl(s) != -1 || suffixPos_iz(s) != -1 {
			s = append(s, 'e')
			goto step1c
		}
		if endsWithDouble(s) {
			s = s[:len(s)-1]
			goto step1c
		}
		if isShortWord(s) {
			s = append(s, 'e')
		}
		goto step1c
	}

step1c:
	// replace suffix y or Y by i if preceded by a non-vowel which is
	// not the first letter of the word (so cry -> cri, by -> by, say -> say)
	if len(s) > 2 {
		switch s[len(s)-1] {
		case 'y', 'Y':
			if !isVowel[s[len(s)-2]] {
				s[len(s)-1] = 'i'
			}
		}
	}
	goto step2

step2:
	r1, r2 = getR1R2(s)
	// Search for the longest among the following suffixes, and,
	// if found and in R1, perform the action indicated

	if found, idx, i := step2Tree.Find(s); found {
		if i >= r1 {
			s = append(s[:i], step2Reps[idx]...)
		}
		goto step3
	}
	if i := suffixPos_ogi(s); i != -1 && i >= r1 {
		if s[i-1] == 'l' {
			s = append(s[:i], 'o', 'g')
		}
		goto step3
	}
	if i := suffixPos_li(s); i != -1 && i >= r1 {
		if isEnding_li[s[i-1]] {
			s = s[:i]
		}
	}

step3:
	r1, r2 = getR1R2(s)
	if found, idx, i := step3Tree.Find(s); found {
		if i >= r1 {
			s = append(s[:i], step3Reps[idx]...)
		}
		goto step4
	}
	if i := suffixPos_ative(s); i != -1 && i >= r2 {
		s = s[:i]
		goto step4
	}

step4:
	r1, r2 = getR1R2(s)
	if found, _, i := step4Tree.Find(s); found {
		if i >= r2 {
			s = s[:i]
		}
		goto step5
	}
	if i := suffixPos_ion(s); i != -1 && i >= r2 {
		switch s[i-1] {
		case 's', 't':
			s = s[:i]
		}
	}

step5:
	r1, r2 = getR1R2(s)
	i := len(s) - 1
	if i > 0 && s[i] == 'e' {
		if i >= r2 {
			s = s[:i]
			goto final
		}

		if i >= r1 {
			// if not preceded by a short syllable
			if i < 3 {
				goto final
			}
			// N + v + N
			last := s[i-1]
			if !isVowel[s[i-3]] && isVowel[s[i-2]] && !isVowel[last] &&
				last != 'w' && last != 'x' && last != 'Y' {
				goto final
			}
			s = s[:i]
		}
		goto final
	}

	if i > 1 && i >= r2 && s[i] == 'l' && s[i-1] == 'l' {
		s = s[:i]
	}

final:
	for i, v := range s {
		if v == 'Y' {
			s[i] = 'y'
		}
	}
	return s
}

var isEnding_li = [256]bool{
	// valid li-ending: c   d   e   g   h   k   m   n   r   t
	'c': true,
	'd': true,
	'e': true,
	'g': true,
	'h': true,
	'k': true,
	'm': true,
	'n': true,
	'r': true,
	't': true,
}

var isEnding_dbl = [256]bool{
	'b': true,
	'd': true,
	'f': true,
	'g': true,
	'm': true,
	'n': true,
	'p': true,
	'r': true,
	't': true,
}

var isVowel = [256]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'y': true,
}

func hasVowelBeforePos(s []byte, pos int) bool {
	for i := pos; i >= 0; i-- {
		if isVowel[s[i]] {
			return true
		}
	}
	return false
}

func calcR(s []byte) int {
	for i := 0; i < len(s)-1; i++ {
		if isVowel[s[i]] && !isVowel[s[i+1]] {
			return i + 2
		}
	}
	return len(s)
}

func getR1(s []byte) (r1 int) {
	n := findRException(s)
	if n >= 0 {
		return n
	}
	return calcR(s)
}

func getR1R2(s []byte) (r1, r2 int) {
	n := findRException(s)
	if n >= 0 {
		return n, n + calcR(s[n:])
	}
	r1 = calcR(s)
	r2 = r1 + calcR(s[r1:])
	return
}

func endsWithDouble(s []byte) bool {
	if len(s) < 2 {
		return false
	}
	last := s[len(s)-1]
	return isEnding_dbl[last] && s[len(s)-2] == last
}

func isShortWord(s []byte) bool {
	if r1 := getR1(s); r1 != len(s) {
		return false
	}
	i := len(s)
	if i == 2 && isVowel[s[0]] && !isVowel[s[1]] {
		return true
	}
	if i < 3 {
		return false
	}
	// ends with short sillable?
	// N + v + N
	last := s[i-1]
	if !isVowel[s[i-3]] && isVowel[s[i-2]] && !isVowel[last] &&
		last != 'w' && last != 'x' && last != 'Y' {
		return true
	}
	return false
}
