package porter2

func findRException(b []byte) int {
	if len(b) < 5 {
		return -1
	}

	// gener, arsen
	if (b[0] == 'g' && b[1] == 'e' && b[2] == 'n' && b[3] == 'e' && b[4] == 'r') ||
		(b[0] == 'a' && b[1] == 'r' && b[2] == 's' && b[3] == 'e' && b[4] == 'n') {
		return 5
	}

	// commun
	if len(b) < 6 {
		return -1
	}

	if b[0] == 'c' && b[1] == 'o' && b[2] == 'm' && b[3] == 'm' && b[4] == 'u' && b[5] == 'n' {
		return 6
	}

	return -1
}

var step1bWords = [][]byte{
	[]byte("ingly"),
	[]byte("edly"),
	[]byte("ing"),
	[]byte("ed"),
}

var step1bTree = buildRevTree(step1bWords, false)

var step2Words = [][]byte{
	[]byte("fulness"), // ful
	[]byte("ousness"), // ous
	[]byte("iveness"), // ive
	[]byte("ational"), // ate
	[]byte("ization"), // ize
	[]byte("tional"),  // tion
	[]byte("biliti"),  // ble
	[]byte("lessli"),  // less
	[]byte("fulli"),   // ful
	[]byte("ousli"),   // ous
	[]byte("iviti"),   // ive
	[]byte("alism"),   // al
	[]byte("ation"),   // ate
	[]byte("entli"),   // ent
	[]byte("aliti"),   // al
	[]byte("enci"),    // ence
	[]byte("anci"),    // ance
	[]byte("abli"),    // able
	[]byte("izer"),    // ize
	[]byte("ator"),    // ate
	[]byte("alli"),    // al
	[]byte("bli"),     // ble
	//"ogi",   // replace with og if preceded by l -- handled later in code
	//"li"     // delete if preceded by a valid li-ending  -- handled later code
}

var step2Tree = buildRevTree(step2Words, false)

var step2Reps = [][]byte{
	[]byte("ful"),
	[]byte("ous"),
	[]byte("ive"),
	[]byte("ate"),
	[]byte("ize"),
	[]byte("tion"),
	[]byte("ble"),
	[]byte("less"),
	[]byte("ful"),
	[]byte("ous"),
	[]byte("ive"),
	[]byte("al"),
	[]byte("ate"),
	[]byte("ent"),
	[]byte("al"),
	[]byte("ence"),
	[]byte("ance"),
	[]byte("able"),
	[]byte("ize"),
	[]byte("ate"),
	[]byte("al"),
	[]byte("ble"),
	//"og"  -- handled later in code
	// ""   -- handled later in code
}

var step3Words = [][]byte{
	[]byte("ational"), // ate
	[]byte("tional"),  // tion
	[]byte("alize"),   // al
	[]byte("icate"),   // ic
	[]byte("iciti"),   // ic
	[]byte("ical"),    // ic
	[]byte("ful"),     // (delete)
	[]byte("ness"),    // (delete)
	//ative -- handled later in code
}

var step3Tree = buildRevTree(step3Words, false)

var step3Reps = [][]byte{
	[]byte("ate"),
	[]byte("tion"),
	[]byte("al"),
	[]byte("ic"),
	[]byte("ic"),
	[]byte("ic"),
	[]byte{},
	[]byte{},
	[]byte{},
}

var step4Words = [][]byte{
	[]byte("ement"),
	[]byte("able"),
	[]byte("ible"),
	[]byte("ance"),
	[]byte("ence"),
	[]byte("ment"),
	[]byte("ant"),
	[]byte("ent"),
	[]byte("ism"),
	[]byte("ate"),
	[]byte("iti"),
	[]byte("ous"),
	[]byte("ive"),
	[]byte("ize"),
	[]byte("al"),
	[]byte("er"),
	[]byte("ic"),
	// "ion" -- delete if preceded by s or t
}

var step4Tree = buildRevTree(step4Words, false)

var exceptions1 = buildMapping([]kv{
	// special changes
	{k: []byte("skis"), v: []byte("ski")},
	{k: []byte("skies"), v: []byte("sky")},
	{k: []byte("dying"), v: []byte("die")},
	{k: []byte("lying"), v: []byte("lie")},
	{k: []byte("tying"), v: []byte("tie")},

	// special -LY cases
	{k: []byte("idly"), v: []byte("idl")},
	{k: []byte("gently"), v: []byte("gentl")},
	{k: []byte("ugly"), v: []byte("ugli")},
	{k: []byte("early"), v: []byte("earli")},
	{k: []byte("only"), v: []byte("onli")},
	{k: []byte("singly"), v: []byte("singl")},
	//invariant forms
	{k: []byte("sky"), v: []byte("sky")},
	{k: []byte("news"), v: []byte("news")},
	{k: []byte("howe"), v: []byte("howe")},
	// not plural forms
	{k: []byte("atlas"), v: []byte("atlas")},
	{k: []byte("cosmos"), v: []byte("cosmos")},
	{k: []byte("bias"), v: []byte("bias")},
	{k: []byte("andes"), v: []byte("andes")},
})

var exceptions2 = buildTree([][]byte{
	[]byte("inning"),
	[]byte("outing"),
	[]byte("canning"),
	[]byte("herring"),
	[]byte("earring"),
	[]byte("proceed"),
	[]byte("exceed"),
	[]byte("succeed"),
}, true)

func isException2(s []byte) bool {
	// []byte("canning"),
	// []byte("earring"),
	// []byte("exceed"),
	// []byte("herring"),
	// []byte("inning"),
	// []byte("outing"),
	// []byte("proceed"),
	// []byte("succeed"),

	if len(s) < 6 || len(s) > 7 {
		return false
	}

	switch s[0] {
	case 'c':
		return len(s) == 7 && s[1] == 'a' && s[2] == 'n' && s[3] == 'n' && s[4] == 'i' && s[5] == 'n' && s[6] == 'g'
	case 'e':
		switch s[1] {
		case 'a':
			return len(s) == 7 && s[2] == 'r' && s[3] == 'r' && s[4] == 'i' && s[5] == 'n' && s[6] == 'g'
		case 'x':
			return len(s) == 6 && s[2] == 'c' && s[3] == 'e' && s[4] == 'e' && s[5] == 'd'
		}
	case 'h':
		return len(s) == 7 && s[1] == 'e' && s[2] == 'r' && s[3] == 'r' && s[4] == 'i' && s[5] == 'n' && s[6] == 'g'
	case 'i':
		return len(s) == 6 && s[1] == 'n' && s[2] == 'n' && s[3] == 'i' && s[4] == 'n' && s[5] == 'g'
	case 'o':
		return len(s) == 6 && s[1] == 'u' && s[2] == 't' && s[3] == 'i' && s[4] == 'n' && s[5] == 'g'
	case 'p':
		return len(s) == 7 && s[1] == 'r' && s[2] == 'o' && s[3] == 'c' && s[4] == 'e' && s[5] == 'e' && s[6] == 'd'
	case 's':
		return len(s) == 7 && s[1] == 'u' && s[2] == 'c' && s[3] == 'c' && s[4] == 'e' && s[5] == 'e' && s[6] == 'd'
	}
	return false
}
