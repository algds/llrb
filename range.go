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
	var result []Value
	if left := recRange(n.Left, cmp, start, end); left != nil {
		result = append(result, left...)
	}
	if cmp(n.K, start) >= 0 && cmp(n.K, end) <= 0 {
		result = append(result, n.V)
	}
	if right := recRange(n.Right, cmp, start, end); right != nil {
		result = append(result, right...)
	}
	return result
}
