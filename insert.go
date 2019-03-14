package llrb

// Insert puts a key-value pair in the tree.
func Insert(t Tree, key Key, val Value) {
	pt := t.(*tree)
	pt.size++
	pt.root = insert(pt, pt.root, pt.compare, key, val)
	pt.root.IsRed = false
}

func insert(pt *tree, n *node, cmp func(a, b Key) int, k Key, v Value) *node {
	if n == nil {
		return &node{K: k, V: v, IsRed: true}
	}
	if isRed(n.Left) && isRed(n.Right) {
		flipColors(n)
	}
	switch c := cmp(k, n.K); {
	case c == 0:
		n.V = v
		pt.size-- // cancel the size increase for an update.
	case c < 0:
		n.Left = insert(pt, n.Left, cmp, k, v)
	case c > 0:
		n.Right = insert(pt, n.Right, cmp, k, v)
	}
	if isRed(n.Right) && !isRed(n.Left) {
		n = rotateLeft(n)
	}
	if isRed(n.Left) && isRed(n.Left.Left) {
		n = rotateRight(n)
	}
	return n
}

func rotateLeft(n *node) *node {
	x := n.Right
	n.Right = x.Left
	x.Left = n
	x.IsRed = n.IsRed
	n.IsRed = true
	return x
}

func rotateRight(n *node) *node {
	x := n.Left
	n.Left = x.Right
	x.Right = n
	x.IsRed = n.IsRed
	n.IsRed = true
	return x
}

func flipColors(n *node) {
	n.IsRed = !n.IsRed
	if n.Left != nil {
		n.Left.IsRed = !n.Left.IsRed
	}
	if n.Right != nil {
		n.Right.IsRed = !n.Right.IsRed
	}
}
