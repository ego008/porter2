package porter2

type tree struct {
	next  [256]*tree
	match bool
	exact bool
}

func buildTree(items [][]byte, exact bool) *tree {
	node := &tree{exact: exact}
	for _, item := range items {
		cur := node
		for _, b := range item {
			if cur.next[b] == nil {
				cur.next[b] = &tree{}
			}
			cur = cur.next[b]
		}
		cur.match = true
	}
	return node
}

func (node *tree) Find(item []byte) (found bool, n int) {
	cur := node
	for i := 0; i < len(item); i++ {
		b := item[i]
		if cur.next[b] == nil {
			break
		}
		cur = cur.next[b]
		if cur.match {
			found = true
			n = i + 1
		}
	}
	if !found || (node.exact && n != len(item)) {
		return false, 0
	}
	return found, n
}

type mapping struct {
	next  [256]*mapping
	match bool
	v     []byte
}

type kv struct{ k, v []byte }

func buildMapping(kvs []kv) *mapping {
	node := &mapping{}
	for _, msg := range kvs {
		cur := node
		for _, b := range msg.k {
			if cur.next[b] == nil {
				cur.next[b] = &mapping{}
			}
			cur = cur.next[b]
		}
		cur.match = true
		cur.v = msg.v
	}
	return node
}

func (node *mapping) Find(s []byte) (v []byte, found bool) {
	cur := node
	var n = 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if cur.next[b] == nil {
			return nil, false
		}
		cur = cur.next[b]
		if cur.match {
			found = true
			v = cur.v
			n = i + 1
		}
	}
	if !found || n != len(s) {
		return nil, false
	}
	return v, true
}

type revTree struct {
	next  [256]*revTree
	idx   int
	match bool
	exact bool
}

func buildRevTree(items [][]byte, exact bool) *revTree {
	node := &revTree{exact: exact}
	for idx, item := range items {
		cur := node
		for i := len(item) - 1; i >= 0; i-- {
			b := item[i]
			if cur.next[b] == nil {
				cur.next[b] = &revTree{}
			}
			cur = cur.next[b]
		}
		cur.idx = idx
		cur.match = true
	}
	return node
}

func (node *revTree) Find(item []byte) (found bool, idx, n int) {
	cur := node
	for i := len(item) - 1; i >= 0; i-- {
		b := item[i]
		if cur.next[b] == nil {
			break
		}
		cur = cur.next[b]
		if cur.match {
			found = true
			idx = cur.idx
			n = i
		}
	}
	if !found || (node.exact && n != len(item)-1) {
		return false, -1, 0
	}
	return found, idx, n
}
