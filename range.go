package llrb

// Range returns all the values between start and end inclusively.
func Range(t Tree, start, end Key) []Value {
	tr := t.(*tree)
	return recRange(tr.root, tr.compare, start, end)
}

func recRange(n *node, cmp func(a, b Key) int, start, end Key) []Value {
	if n == nil {
		return nil
	}
	if c := cmp(n.K, start); c < 0 {
		return recRange(n.Right, cmp, start, end)
	} else if c == 0 {
		return append([]Value{n.V}, recRange(n.Right, cmp, start, end)...)
	}
	if c := cmp(n.K, end); c > 0 {
		return recRange(n.Left, cmp, start, end)
	} else if c == 0 {
		return append(recRange(n.Left, cmp, start, end), n.V)
	}
	result := append(recRange(n.Left, cmp, start, end), n.V)
	return append(result, recRange(n.Right, cmp, start, end)...)
}
