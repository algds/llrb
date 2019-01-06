package llrb

// Search attempts to find the Key k in the Tree.
// If found, the Value and true are returned.
// If not found, nil and false are returned.
func Search(t Tree, k Key) (Value, bool) {
	tr := t.(*tree)
	return search(tr.root, tr.compare, k)
}

func search(n *node, cmp func(a, b Key) int, k Key) (Value, bool) {
	if n == nil {
		return nil, false
	}
	if c := cmp(n.K, k); c == 0 {
		return n.V, true
	} else if c > 0 {
		return search(n.Left, cmp, k)
	}
	return search(n.Right, cmp, k)
}
