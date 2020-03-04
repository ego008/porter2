package porter2

// FIXME: IsInBounds abounds

func suffixPos_s(s []byte) int {
	l := len(s)
	if l >= 1 && s[l-1] == 's' {
		return len(s) - 1
	}
	return -1
}

func suffixPos_at(s []byte) int {
	l := len(s)
	if l >= 2 && s[l-2] == 'a' && s[l-1] == 't' {
		return len(s) - 2
	}
	return -1
}

func suffixPos_bl(s []byte) int {
	l := len(s)
	if l >= 2 && s[l-2] == 'b' && s[l-1] == 'l' {
		return len(s) - 2
	}
	return -1
}

func suffixPos_iz(s []byte) int {
	l := len(s)
	if l >= 2 && s[l-2] == 'i' && s[l-1] == 'z' {
		return len(s) - 2
	}
	return -1
}

func suffixPos_li(s []byte) int {
	l := len(s)
	if l >= 2 && s[l-2] == 'l' && s[l-1] == 'i' {
		return len(s) - 2
	}
	return -1
}

func suffixPos_ss(s []byte) int {
	l := len(s)
	if l >= 2 && s[l-2] == 's' && s[l-1] == 's' {
		return len(s) - 2
	}
	return -1
}

func suffixPos_us(s []byte) int {
	l := len(s)
	if l >= 2 && s[l-2] == 'u' && s[l-1] == 's' {
		return len(s) - 2
	}
	return -1
}

func suffixPos_eed(s []byte) int {
	l := len(s)
	if l >= 3 && s[l-3] == 'e' && s[l-2] == 'e' && s[l-1] == 'd' {
		return len(s) - 3
	}
	return -1
}

func suffixPos_ied(s []byte) int {
	l := len(s)
	if l >= 3 && s[l-3] == 'i' && s[l-2] == 'e' && s[l-1] == 'd' {
		return len(s) - 3
	}
	return -1
}

func suffixPos_ies(s []byte) int {
	l := len(s)
	if l >= 3 && s[l-3] == 'i' && s[l-2] == 'e' && s[l-1] == 's' {
		return len(s) - 3
	}
	return -1
}

func suffixPos_ion(s []byte) int {
	l := len(s)
	if l >= 3 && s[l-3] == 'i' && s[l-2] == 'o' && s[l-1] == 'n' {
		return len(s) - 3
	}
	return -1
}

func suffixPos_ogi(s []byte) int {
	l := len(s)
	if l >= 3 && s[l-3] == 'o' && s[l-2] == 'g' && s[l-1] == 'i' {
		return len(s) - 3
	}
	return -1
}

func suffixPos_sses(s []byte) int {
	l := len(s)
	if l >= 4 && s[l-4] == 's' && s[l-3] == 's' && s[l-2] == 'e' && s[l-1] == 's' {
		return len(s) - 4
	}
	return -1
}

func suffixPos_ative(s []byte) int {
	l := len(s)
	if l >= 5 && s[l-5] == 'a' && s[l-4] == 't' && s[l-3] == 'i' && s[l-2] == 'v' && s[l-1] == 'e' {
		return len(s) - 5
	}
	return -1
}

func suffixPos_eedly(s []byte) int {
	l := len(s)
	if l >= 5 && s[l-5] == 'e' && s[l-4] == 'e' && s[l-3] == 'd' && s[l-2] == 'l' && s[l-1] == 'y' {
		return len(s) - 5
	}
	return -1
}

func removeSuffix_apos_s_apos(s []byte) []byte {
	l := len(s)
	if l >= 3 && s[l-3] == '\'' && s[l-2] == 's' && s[l-1] == '\'' {
		return s[:l-3]
	}
	return s
}

func removeSuffix_apos_s(s []byte) []byte {
	l := len(s)
	if l >= 2 && s[l-2] == '\'' && s[l-1] == 's' {
		return s[:l-2]
	}
	return s
}

func removeSuffix_apos(s []byte) []byte {
	l := len(s)
	if l >= 1 && s[l-1] == '\'' {
		return s[:l-1]
	}
	return s
}
