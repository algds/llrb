package llrb

// Tree is a left-leaning red-black tree.
type Tree interface {
	Size() int
	compare(a, b Key) int
}

// New returns a new left-leaning red-black tree with a given
// compare function. The Key type accepted by the compare function is
// the only Key type that will work during an Insert.
// If nil is passed in, then the default integer compare function is used.
func New(cmpFunc func(a, b Key) int) Tree {
	if cmpFunc == nil {
		cmpFunc = func(a, b Key) int {
			ai := a.(int)
			bi := b.(int)
			if ai < bi {
				return -1
			} else if ai == bi {
				return 0
			}
			return 1
		}
	}
	return &tree{nil, 0, cmpFunc}
}

type tree struct {
	root *node
	size int
	// Compare compares two keys. -1 if a<b, 0 if a==b, and 1 if a>b
	cmpFunc func(a, b Key) int
}

// Size returns the number of elements in the tree.
func (t *tree) Size() int {
	return t.size
}

func (t *tree) compare(a, b Key) int {
	return t.cmpFunc(a, b)
}

// Key represents any comparable value.
// https://golang.org/ref/spec#Comparison_operators
type Key interface{}

// Value represents any kind of value that is associated with the key.
type Value interface{}

type node struct {
	K           Key
	V           Value
	Left, Right *node
	IsRed       bool
}

func isRed(n *node) bool {
	return n != nil && n.IsRed
}
