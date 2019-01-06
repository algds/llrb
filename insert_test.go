package llrb

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t.Parallel()
	tr := New(nil)
	Insert(tr, 5, "foo")
	Insert(tr, 10, "baz")
	Insert(tr, 1, "boo")
	Insert(tr, 2, "boo")
	Insert(tr, 3, "boo")
	Insert(tr, 4, "boo")
	Insert(tr, 8, "boo")
	Insert(tr, 7, "boo")
	Insert(tr, 6, "boo")
	if tr.Size() != 9 {
		t.Errorf("Size doesn't equal 9")
	}
	Insert(tr, 10, "just a value update")
	if tr.Size() != 9 {
		t.Errorf("Inserting the same key twice is a value update only.")
	}
	if tr.Size() != walkTreeSize(tr.(*tree).root) {
		t.Errorf("Size should equal actual size")
	}
}

func walkTreeSize(n *node) int {
	if n == nil {
		return 0
	}
	return walkTreeSize(n.Left) + walkTreeSize(n.Right) + 1
}
