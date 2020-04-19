package porter2

// FIXME: IsInBounds abounds

func suffixHas_at_bl_iz(s []byte) bool {
	off := len(s) - 2
	if off < 0 {
		return false
	}
	s = s[off:]
	_ = s[1]
	return (s[0] == 'a' && s[1] == 't') ||
		(s[0] == 'b' && s[1] == 'l') ||
		(s[0] == 'i' && s[1] == 'z')
}

func suffixHas_ss_us(s []byte) bool {
	off := len(s) - 2
	if off < 0 {
		return false
	}
	s = s[off:]
	_ = s[1]
	return s[1] == 's' && (s[0] == 'u' || s[0] == 's')
}

func suffixPos_s(s []byte) int {
	last := len(s) - 1
	if last >= 0 && s[last] == 's' {
		return last
	}
	return -1
}

func suffixPos_li(s []byte) int {
	off := len(s) - 2
	if off >= 0 {
		s = s[off:]
		_ = s[1]
		if s[0] == 'l' && s[1] == 'i' {
			return off
		}
	}
	return -1
}

func suffixPos_eed(s []byte) int {
	off := len(s) - 3
	if off >= 0 {
		s = s[off:]
		_ = s[2]
		if s[0] == 'e' && s[1] == 'e' && s[2] == 'd' {
			return off
		}
	}
	return -1
}

func suffixPos_ied_ies(s []byte) int {
	off := len(s) - 3
	if off >= 0 {
		s = s[off:]
		_ = s[2]
		if s[0] == 'i' && s[1] == 'e' && (s[2] == 'd' || s[2] == 's') {
			return off
		}
	}
	return -1
}

func suffixPos_ion(s []byte) int {
	off := len(s) - 3
	if off >= 0 {
		s = s[off:]
		_ = s[2]
		if s[0] == 'i' && s[1] == 'o' && s[2] == 'n' {
			return off
		}
	}
	return -1
}

func suffixPos_ogi(s []byte) int {
	off := len(s) - 3
	if off >= 0 {
		s = s[off:]
		_ = s[2]
		if s[0] == 'o' && s[1] == 'g' && s[2] == 'i' {
			return off
		}
	}
	return -1
}

func suffixPos_sses(s []byte) int {
	off := len(s) - 4
	if off >= 0 {
		s = s[off:]
		_ = s[3]
		if s[0] == 's' && s[1] == 's' && s[2] == 'e' && s[3] == 's' {
			return off
		}
	}
	return -1
}

func suffixPos_ative(s []byte) int {
	off := len(s) - 5
	if off >= 0 {
		s = s[off:]
		_ = s[4]
		if s[0] == 'a' && s[1] == 't' && s[2] == 'i' && s[3] == 'v' && s[4] == 'e' {
			return off
		}
	}
	return -1
}

func suffixPos_eedly(s []byte) int {
	off := len(s) - 5
	if off >= 0 {
		s = s[off:]
		_ = s[4]
		if s[0] == 'e' && s[1] == 'e' && s[2] == 'd' && s[3] == 'l' && s[4] == 'y' {
			return off
		}
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
